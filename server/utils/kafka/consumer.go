package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BharathKumarRavichandran/k8s-playground/server/db"
	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"github.com/gocql/gocql"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func initConsumer(kafkaConfig utils.KafkaConfig) {

	topics := []string{kafkaConfig.Topic}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     kafkaConfig.ServiceName,
		"broker.address.family": "v4",
		"group.id":              kafkaConfig.ConsumerGroup,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
		"sasl.mechanisms":       kafkaConfig.SaslMechanisms,
		"security.protocol":     kafkaConfig.SecurityProtocol,
		"sasl.username":         kafkaConfig.Username,
		"sasl.password":         kafkaConfig.Password,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
	}

	utils.Logger.Infof("Created Consumer %v\n", c)

	err = c.SubscribeTopics(topics, nil)

	// Create a new go-routine for the consumer
	run := true
	go func() {
		for run == true {
			select {

			case sig := <-sigchan:
				utils.Logger.Infof("Caught signal %v: terminating\n", sig)
				run = false

				utils.Logger.Info("Closing consumer\n")
				c.Close()

			default:
				ev := c.Poll(100)
				if ev == nil {
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:

					// Push Message to database
					message := string(e.Value)
					if err := db.Session.Query(`INSERT INTO records (id, message, created_date) VALUES (?, ?, ?)`,
						gocql.TimeUUID(), message, time.Now()).Exec(); err != nil { // can also use gocql.RandomUUID()
						utils.Logger.Error(err)
					}

					utils.Logger.Infof("%% Message on %s:\n%s\n",
						e.TopicPartition, string(e.Value))
					if e.Headers != nil {
						utils.Logger.Infof("%% Headers: %v\n", e.Headers)
					}
				case kafka.Error:
					// Errors should generally be considered
					// informational, the client will try to
					// automatically recover.
					// Here we choose to terminate
					// the application if all brokers are down.
					//fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
					if e.Code() == kafka.ErrAllBrokersDown {
						run = false
					}
				default:
					utils.Logger.Infof("Ignored %v\n", e)
				}
			}
		}
	}()

}
