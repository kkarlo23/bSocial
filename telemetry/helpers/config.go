package helpers

import (
	"encoding/json"
	"os"
)

type KafkaTopic struct {
	Name            string `json:"name"`
	ConsumerGroupID string `json:"consumerGroupID"`
	Partition       int    `json:"partition"`
}
type KafkaConfig struct {
	KafkaProtocol     string     `json:"kafkaProtocol"`
	KafkaHost         string     `json:"kafkaHost"`
	KafkaPort         string     `json:"kafkaPort"`
	UserRegisterTopic KafkaTopic `json:"userRegisterTopic"`
	PostTopic         KafkaTopic `json:"postTopic"`
	CommentTopic      KafkaTopic `json:"commentTopic"`
}

type Index struct {
	UserRegister string `json:"userRegister"`
	Post         string `json:"post"`
	Comment      string `json:"comment"`
}

type ElasticSearchConfig struct {
	ElasticSearchProtocol string `json:"elasticSearchProtocol"`
	ElasticSearchHost     string `json:"elasticSearchHost"`
	ElasticSearchPort     string `json:"elasticSearchPort"`
	Indices               Index  `json:"indices"`
}

type Config struct {
	Kafka         KafkaConfig         `json:"kafka"`
	ElasticSearch ElasticSearchConfig `json:"elasticSearch"`
}

var CONFIG Config

// unpacks config file
func InitConfig() error {
	jsonData, err := os.ReadFile("./config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &CONFIG)
	if err != nil {
		return err
	}

	return nil
}
