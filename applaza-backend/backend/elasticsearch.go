package backend

import (
	"applaza-backend/constants"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var (
	ESBackend *ElasticsearchBackend
)

type ElasticsearchBackend struct {
	client *elastic.Client
}

func InitElasticsearchBackend() {
	client, err := elastic.NewClient(
		elastic.SetURL(constants.ES_URL),
		elastic.SetBasicAuth(constants.ES_USERNAME, constants.ES_PASSWORD),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	exists, err := client.IndexExists(constants.APP_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		mapping := `{
						"mappings": {
							"properties": {
								"id": {"type": "keyword"},
								"user": {"type": "keyword"},
								"title": {"type": "text"},
								"description": {"type": "text"},
								"price": {"type": "keyword", "index": false},
								"url": {"type": "keyword", "index": false}
							}
						}
					}`
		_, err := client.CreateIndex(constants.APP_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Index is created.")

	ESBackend = &ElasticsearchBackend{client: client}
}

func (backend *ElasticsearchBackend) ReadFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	searchResult, err := backend.client.Search().Index(index).Query(query).Pretty(true).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return searchResult, nil
}
