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

func ConsumePost(wg *sync.WaitGroup) {
	r := CreateReader(helpers.CONFIG.Kafka.PostTopic)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var post domain.KafkaPost
		err = json.Unmarshal(m.Value, &post)
		if err != nil {
			fmt.Println("could not unmarshal post")
		}

		doc, err := elasticSearch.CreatePost(post)
		if err != nil {
			fmt.Println("errored", err)
		}

		fmt.Printf("post document created, with ID: %s\n", doc.Id)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
		wg.Done()
	}
}
