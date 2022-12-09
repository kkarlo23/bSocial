package kafkaProducer

import (
	"bSocial/domain"
	"bSocial/helpers"
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func ProduceComment(commentData domain.KafkaComment) error {
	topic := helpers.CONFIG.Kafka.CommentTopic.Name
	partition := helpers.CONFIG.Kafka.CommentTopic.Partition

	conn, err := kafka.DialLeader(context.Background(),
		helpers.CONFIG.Kafka.KafkaProtocol,
		fmt.Sprintf("%s:%s", helpers.CONFIG.Kafka.KafkaHost, helpers.CONFIG.Kafka.KafkaPort),
		topic,
		partition)

	if err != nil {
		return err
	}
	postJson, err := json.Marshal(commentData)
	if err != nil {
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: postJson},
	)

	if err != nil {
		return err
	}

	if err := conn.Close(); err != nil {
		return err
	}
	return nil
}
