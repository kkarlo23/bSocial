package kafkaConsumer

import (
	"fmt"
	"telemetry/helpers"

	kafka "github.com/segmentio/kafka-go"
)

func CreateReader(topic helpers.KafkaTopic) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s",
			helpers.CONFIG.Kafka.KafkaHost,
			helpers.CONFIG.Kafka.KafkaPort)},
		GroupID:  topic.ConsumerGroupID,
		Topic:    topic.Name,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
