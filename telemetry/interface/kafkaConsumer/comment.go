package kafkaConsumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"telemetry/domain"
	"telemetry/helpers"
	"telemetry/interface/elasticSearch"
)

func ConsumeComment(wg *sync.WaitGroup) {
	r := CreateReader(helpers.CONFIG.Kafka.CommentTopic)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var comment domain.KafkaComment
		err = json.Unmarshal(m.Value, &comment)
		if err != nil {
			fmt.Println("could not unmarshal comment")
		}

		doc, err := elasticSearch.CreateComment(comment)
		if err != nil {
			fmt.Println("errored", err)
		}

		fmt.Printf("comment document created, with ID: %s\n", doc.Id)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
		wg.Done()
	}
}
