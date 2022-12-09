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

func ProducePost(postData domain.KafkaPost) error {
	topic := helpers.CONFIG.Kafka.PostTopic.Name
	partition := helpers.CONFIG.Kafka.PostTopic.Partition

	conn, err := kafka.DialLeader(context.Background(),
		helpers.CONFIG.Kafka.KafkaProtocol,
		fmt.Sprintf("%s:%s", helpers.CONFIG.Kafka.KafkaHost, helpers.CONFIG.Kafka.KafkaPort),
		topic,
		partition)

	if err != nil {
		return err
	}
	postJson, err := json.Marshal(postData)
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
