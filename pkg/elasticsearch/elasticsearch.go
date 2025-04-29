package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"main/pkg/db"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/lib/pq"
)

var ESClient *elasticsearch.Client

type FilmDTO struct {
	FilmID          int            `json:"film_id"`
	Title           string         `json:"title"`
	Description     *string        `json:"description"`
	ReleaseYear     int            `json:"release_year"`
	LanguageID      uint16         `json:"language_id"`
	RentalDuration  uint16         `json:"rental_duration"`
	RentalRate      float32        `json:"rental_rate"`
	Length          uint16         `json:"length"`
	ReplacementCost float32        `json:"replacement_cost"`
	Rating          string         `json:"rating"`
	LastUpdate      time.Time      `json:"last_update"`
	SpecialFeatures pq.StringArray `json:"special_features" gorm:"type:text[]"`
	Fulltext        string         `json:"fulltext"`
}

type FilmFilterDTO struct {
	Title       string  `form:"title"`
	Description *string `form:"description"`
	ReleaseYear int     `form:"release_year"`
}

func (d FilmDTO) ToFilm() FilmDTO {
	return FilmDTO{
		FilmID:          d.FilmID,
		Title:           d.Title,
		Description:     d.Description,
		ReleaseYear:     d.ReleaseYear,
		LanguageID:      d.LanguageID,
		RentalDuration:  d.RentalDuration,
		RentalRate:      d.RentalRate,
		Length:          d.Length,
		ReplacementCost: d.ReplacementCost,
		Rating:          d.Rating,
		LastUpdate:      d.LastUpdate,
		SpecialFeatures: d.SpecialFeatures,
		Fulltext:        d.Fulltext,
	}
}

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

func IndexFilmsToES(films []FilmDTO) error {
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

func SearchFilms(pagination db.Pagination, incomingFilters []byte) ([]byte, int64, error) {
	var filters FilmFilterDTO
	err := json.Unmarshal(incomingFilters, &filters)
	if err != nil {
		fmt.Printf("Error while serializing:: %v", err)
	}

	must := []map[string]interface{}{}

	if filters.Title != "" {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"title": filters.Title,
			},
		})
	}
	if filters.Description != nil {
		must = append(must, map[string]interface{}{
			"match": map[string]interface{}{
				"description": *filters.Description,
			},
		})
	}
	if filters.ReleaseYear != 0 {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				"release_year": filters.ReleaseYear,
			},
		})
	}

	queryBody := map[string]interface{}{
		"from": pagination.GetOffset(),
		"size": pagination.GetLimit(),
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
			},
		},
	}

	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, err
	}

	req := esapi.SearchRequest{
		Index: []string{"films"},
		Body:  bytes.NewReader(body),
	}

	res, err := req.Do(context.Background(), ESClient)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	var r struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source FilmDTO `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	var films []FilmDTO
	for _, hit := range r.Hits.Hits {
		films = append(films, hit.Source.ToFilm())
	}

	filmsJSON, err := json.Marshal(films)
	if err != nil {
		return nil, 0, err
	}

	return filmsJSON, r.Hits.Total.Value, nil
}
