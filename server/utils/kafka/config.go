package kafka

import "github.com/BharathKumarRavichandran/k8s-playground/server/utils"

type KafkaConfig struct {
	serviceName   string
	consumerGroup string
	topic         string
	username      string
	password      string
}

var kafkaConfig KafkaConfig

func Init(config *utils.Config) {

	kafkaConfig.serviceName = config.KAFKA_SERVICE_NAME
	kafkaConfig.consumerGroup = config.KAFKA_CONSUMER_GROUP
	kafkaConfig.topic = config.KAFKA_TOPIC
	kafkaConfig.username = config.KAFKA_USERNAME
	kafkaConfig.password = config.KAFKA_PASSWORD

	initProducer(kafkaConfig)
	initConsumer(kafkaConfig)
}

func GetConfiguration() KafkaConfig {
	return kafkaConfig
}
