package elasticSearch

import (
	"context"
	"encoding/json"
	"telemetry/domain"
	"telemetry/helpers"
	"time"

	"github.com/olivere/elastic/v7"
)

func CreatePost(post domain.KafkaPost) (*elastic.IndexResponse, error) {

	post.CreatedAt = post.CreatedAt.Add(RandDay(1, 20) * 24 * time.Hour)
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
