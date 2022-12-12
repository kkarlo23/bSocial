package elasticSearch

import (
	"context"
	"encoding/json"
	"telemetry/domain"
	"telemetry/helpers"

	"github.com/olivere/elastic/v7"
)

func CreatePost(post domain.KafkaPost) (*elastic.IndexResponse, error) {
	dataJSON, _ := json.Marshal(post)
	jsonString := string(dataJSON)
	res, err := ESClient.Index().
		Index(helpers.CONFIG.ElasticSearch.Indices.Post).
		BodyJson(jsonString).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return res, nil
}
