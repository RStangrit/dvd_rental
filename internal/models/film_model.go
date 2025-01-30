package models

import (
	"time"

	"github.com/lib/pq"
)

type Film struct {
	FilmID          int            `json:"film_id" gorm:"type: integer;primaryKey;autoIncrement, not null"`
	Title           string         `json:"title" gorm:"type: varchar(255) not null"`
	Description     *string        `json:"description" gorm:"type:text"`
	ReleaseYear     int            `json:"release_year" gorm:"type:int4"`
	LanguageID      uint16         `json:"language_id" gorm:"not null;foreignKey:LanguageID"`
	RentalDuration  uint16         `json:"rental_duration" gorm:"type: int2 not null;default:3"`
	RentalRate      float32        `json:"rental_rate" gorm:"type:numeric(4,2);not null;default:4.99"`
	Length          uint16         `json:"length" gorm:"type:smallint"`
	ReplacementCost float32        `json:"replacement_cost" gorm:"type:numeric(5,2);not null;default:19.99"`
	Rating          string         `json:"rating"`
	LastUpdate      time.Time      `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime"`
	SpecialFeatures pq.StringArray `json:"special_features" gorm:"type:text[]"`
	Fulltext        string         `json:"fulltext" gorm:"type:tsvector"`
	FilmsActor      []FilmActor    `json:"-" gorm:"foreignKey:FilmID"`
	FilmsCategories []FilmCategory `json:"-" gorm:"foreignKey:FilmID"`
	FilmsInventory  []Inventory    `json:"-" gorm:"foreignKey:FilmID"`
}

func (Film) TableName() string {
	return "film"
}
