package elasticSearch

import (
	// "github.com/elastic/go-elasticsearch/v7"

	"fmt"
	"telemetry/helpers"

	"github.com/olivere/elastic/v7"
)

// var ESClient *elasticsearch.Client
var ESClient *elastic.Client

func InitESClient() error {

	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("%s://%s:%s",
		helpers.CONFIG.ElasticSearch.ElasticSearchProtocol,
		helpers.CONFIG.ElasticSearch.ElasticSearchHost,
		helpers.CONFIG.ElasticSearch.ElasticSearchPort)))
	if err != nil {
		println(err)
		return err
	}
	ESClient = client

	return nil
}
