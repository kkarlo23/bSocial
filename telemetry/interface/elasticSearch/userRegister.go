package elasticSearch

import (
	"context"
	"encoding/json"
	"math/rand"
	"telemetry/domain"
	"telemetry/helpers"
	"time"

	"github.com/olivere/elastic/v7"
)

func RandDay(min, max int) time.Duration {
	return time.Duration(rand.Intn(max-min)+min) * -1
}
func CreateUserRegister(user domain.User) (*elastic.IndexResponse, error) {
	dataJSON, _ := json.Marshal(user)
	jsonString := string(dataJSON)
	res, err := ESClient.Index().
		Index(helpers.CONFIG.ElasticSearch.Indices.UserRegister).
		BodyJson(jsonString).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return res, nil
}
