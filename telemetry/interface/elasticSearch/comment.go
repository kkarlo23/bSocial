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

	// Update comment count on post when new comment is created
	_, err = ESClient.UpdateByQuery().
		Query(elastic.NewMatchQuery("id", comment.PostID)).
		Index(helpers.CONFIG.ElasticSearch.Indices.Post).
		Script(elastic.NewScript("ctx._source.commentCount+=1")).
		Do(context.Background())

	if err != nil {
		println(err.Error())
	}

	return res, nil

}
