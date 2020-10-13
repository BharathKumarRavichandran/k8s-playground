package kafka

import (
	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var producer *kafka.Producer

func initProducer(kafkaConfig utils.KafkaConfig) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  kafkaConfig.ServiceName,
		"enable.idempotence": true, // Enable the Idempotent Producer
		"sasl.mechanisms":    kafkaConfig.SaslMechanisms,
		"security.protocol":  kafkaConfig.SecurityProtocol,
		"sasl.username":      kafkaConfig.Username,
		"sasl.password":      kafkaConfig.Password,
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
	topic := utils.GetConfiguration().Kafka.Topic
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

	utils.Logger.Info("Closing Kafka-Producer delivery channel...")
	close(deliveryChan)
}

func Close() {
	producer.Close()
	utils.Logger.Infof("Kafka Producer connection closed")
}
