package helpers

import (
	"encoding/json"
	"os"
)

type MySQLConfig struct {
	DbUser     string `json:"dbUser"`
	DbPassword string `json:"dbPassword"`
	DbHost     string `json:"dbHost"`
	DbPort     string `json:"dbPort"`
}

type KafkaTopic struct {
	Name      string `json:"name"`
	Partition int    `json:"partition"`
}
type KafkaConfig struct {
	KafkaProtocol     string     `json:"kafkaProtocol"`
	KafkaHost         string     `json:"kafkaHost"`
	KafkaPort         string     `json:"kafkaPort"`
	UserRegisterTopic KafkaTopic `json:"userRegisterTopic"`
	PostTopic         KafkaTopic `json:"postTopic"`
	CommentTopic      KafkaTopic `json:"commentTopic"`
}
type JsonConfig struct {
	Secret    string `json:"secret"`
	ExpMinute int    `json:"expMinute"`
}

type Config struct {
	MySQL MySQLConfig `json:"mySql"`
	Kafka KafkaConfig `json:"kafka"`
	Json  JsonConfig  `json:"json"`
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
