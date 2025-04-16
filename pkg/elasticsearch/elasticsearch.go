package elasticsearch

import (
	"fmt"

	elasticsearch "github.com/elastic/go-elasticsearch/v9"
)

var ESClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"ELASTICSEARCH_URL",
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("Error creating the client: %s", err)
	}
	ESClient = client
	fmt.Println("Elasticsearch client initialized")
}
