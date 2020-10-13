package kafka

import (
	"fmt"

	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var producer *kafka.Producer

func initProducer(kafkaConfig KafkaConfig) {

	brokers := kafkaConfig.serviceName

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		// Enable the Idempotent Producer
		"enable.idempotence": true,
		"sasl.mechanisms":    "SCRAM-SHA-256",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      kafkaConfig.username,
		"sasl.password":      kafkaConfig.password,
	})

	producer = p

	if err != nil {
		utils.Logger.Fatalf("Failed to create producer: %s\n", err)
	}

	utils.Logger.Infof("Created Producer %v\n", producer)
}

func ProduceMessage(message string) {
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	topic := kafkaConfig.topic
	deliveryChan := make(chan kafka.Event)

	utils.Logger.Infof("Pushing message: %s", message)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "TestHeader", Value: []byte("Header values are binary")}},
	}, deliveryChan)

	if err != nil {
		utils.Logger.Errorf("Failed to push message: %s\n", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		utils.Logger.Errorf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		utils.Logger.Infof("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	fmt.Println("Closing delivery channel...")
	close(deliveryChan)
}

func Close() {
	producer.Close()
	utils.Logger.Infof("Kafka Producer connection closed")
}
