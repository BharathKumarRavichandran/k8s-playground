package kafka

import (
	"context"
	"time"

	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func CreateTopic(topic string) {

	kafkaConfig := utils.GetConfiguration().Kafka

	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers":       kafkaConfig.ServiceName,
		"broker.version.fallback": "0.10.0.0",
		"api.version.fallback.ms": 0,
		"sasl.mechanisms":         kafkaConfig.SaslMechanisms,
		"security.protocol":       kafkaConfig.SecurityProtocol,
		"sasl.username":           kafkaConfig.Username,
		"sasl.password":           kafkaConfig.Password,
	})

	if err != nil {
		utils.Logger.Fatalf("Failed to create Admin client: %s\n", err)
	}

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create topics on cluster.
	// Set Admin options to wait for the operation to finish (or at most 60s)
	maxDuration, err := time.ParseDuration("60s")
	if err != nil {
		panic("time.ParseDuration(60s)")
	}

	results, err := adminClient.CreateTopics(ctx,
		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 3}},
		kafka.SetAdminOperationTimeout(maxDuration))

	if err != nil {
		utils.Logger.Fatalf("Problem during the topic creation: %v\n", err)
	}

	// Check for specific topic errors
	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError &&
			result.Error.Code() != kafka.ErrTopicAlreadyExists {
			utils.Logger.Errorf("Topic creation failed for %s: %v",
				result.Topic, result.Error.String())
		}
	}

	adminClient.Close()
}
