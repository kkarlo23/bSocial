package main

import (
	"log"
	"sync"
	"telemetry/helpers"
	"telemetry/interface/elasticSearch"
	"telemetry/interface/kafkaConsumer"
)

func main() {
	err := helpers.InitConfig()
	if err != nil {
		log.Fatalf("Error initialising config: %s", err)
	}
	err = elasticSearch.InitESClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go kafkaConsumer.ConsumeUserRegister(wg)
	wg.Add(1)
	go kafkaConsumer.ConsumePost(wg)
	wg.Add(1)
	go kafkaConsumer.ConsumeComment(wg)
	wg.Wait()

}
