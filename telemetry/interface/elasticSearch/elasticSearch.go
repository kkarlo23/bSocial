package elasticSearch

import (
	// "github.com/elastic/go-elasticsearch/v7"

	"fmt"
	"telemetry/helpers"

	"github.com/olivere/elastic/v7"
)

// var ESClient *elasticsearch.Client
var ESClient *elastic.Client

// func InitESClient() error {
// 	cfg := elasticsearch.Config{
// 		Addresses: []string{
// 			fmt.Sprintf("%s://%s:%s",
// 				helpers.CONFIG.ElasticSearch.ElasticSearchProtocol,
// 				helpers.CONFIG.ElasticSearch.ElasticSearchHost,
// 				helpers.CONFIG.ElasticSearch.ElasticSearchPort),
// 		},
// 		// ...
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	if err != nil {
// 		return err
// 	}
// 	ESClient = es
// 	return nil
// 	// res, err := es.Info()
// 	// if err != nil {
// 	// 	log.Fatalf("Error getting response: %s", err)
// 	// }

//		// defer res.Body.Close()
//		// log.Println(res)
//	}

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
