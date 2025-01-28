package film

import "time"

type Film struct {
	FilmID          int       `json:"film_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ReleaseYear     int       `json:"release_year"`
	LanguageID      int16     `json:"language_id"`
	RentalDuration  int16     `json:"rental_duration"`
	RentalRate      float64   `json:"rental_rate"`
	Length          int16     `json:"length"`
	ReplacementCost float64   `json:"replacement_cost"`
	Rating          string    `json:"rating"`
	LastUpdate      time.Time `json:"last_update"`
	SpecialFeatures []string  `json:"special_features"`
	Fulltext        string    `json:"fulltext"`
}

func (Film) TableName() string {
	return "film"
}
