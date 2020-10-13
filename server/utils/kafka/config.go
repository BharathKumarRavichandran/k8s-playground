package kafka

import "github.com/BharathKumarRavichandran/k8s-playground/server/utils"

var kafkaConfig *utils.KafkaConfig

func Init(config *utils.Config) {
	initProducer(config.Kafka)
	initConsumer(config.Kafka)
}

func GetConfiguration() *utils.KafkaConfig {
	return kafkaConfig
}
