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
type JsonConfig struct {
	Secret    string `json:"secret"`
	ExpMinute int    `json:"expMinute"`
}

type Config struct {
	MySQL MySQLConfig `json:"mySql"`
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
