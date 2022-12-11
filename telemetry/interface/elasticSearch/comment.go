package elasticSearch

import (
	"context"
	"encoding/json"
	"telemetry/domain"
	"telemetry/helpers"

	"github.com/olivere/elastic/v7"
)

func CreateComment(comment domain.KafkaComment) (*elastic.IndexResponse, error) {
	dataJSON, _ := json.Marshal(comment)
	jsonString := string(dataJSON)
	res, err := ESClient.Index().
		Index(helpers.CONFIG.ElasticSearch.Indices.Comment).
		BodyJson(jsonString).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	return res, nil
}
