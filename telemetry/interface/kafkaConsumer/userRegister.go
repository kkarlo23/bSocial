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

func ConsumeUserRegister(wg *sync.WaitGroup) {
	r := CreateReader(helpers.CONFIG.Kafka.UserRegisterTopic)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var user domain.User
		err = json.Unmarshal(m.Value, &user)
		if err != nil {
			fmt.Println("could not unmarshal user")
		}

		doc, err := elasticSearch.CreateUserRegister(user)
		if err != nil {
			fmt.Println("errored", err)
		}

		fmt.Printf("user-register document created, with ID: %s\n", doc.Id)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
		wg.Done()
	}
}
