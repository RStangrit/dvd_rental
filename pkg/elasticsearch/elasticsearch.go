package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"main/internal/film"

	elasticsearch "github.com/elastic/go-elasticsearch/v9"
)

var ESClient *elasticsearch.Client

func InitElasticsearch() {
	params := config.LoadConfig()
	cfg := elasticsearch.Config{
		Addresses: []string{
			params.ELASTICSEARCH_HOST,
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("Error creating the client: %s", err)
	}
	ESClient = client
	fmt.Println("Elasticsearch client initialized")
}

func IndexFilmsToES(films []film.Film) error {
	for _, film := range films {
		data, err := json.Marshal(film)
		if err != nil {
			return err
		}

		res, err := ESClient.Index(
			"films",
			bytes.NewReader(data),
			ESClient.Index.WithDocumentID(fmt.Sprintf("film-%d", film.FilmID)),
			ESClient.Index.WithRefresh("true"),
		)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.IsError() {
			body, _ := io.ReadAll(res.Body)
			fmt.Printf("Failed to index DVD %d: %s", film.FilmID, string(body))
		}
	}
	return nil
}
