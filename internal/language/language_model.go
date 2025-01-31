package language

import (
	"main/internal/film"
	"main/pkg/db"
	"time"
)

type Language struct {
	LanguageID int         `json:"language_id" gorm:"type: integer;primaryKey;autoIncrement"`
	Name       string      `json:"name" gorm:"type: bpchar(20) not null"`
	LastUpdate time.Time   `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime"`
	Films      []film.Film `json:"-" gorm:"foreignKey:LanguageID"`
}

func (Language) TableName() string {
	return "language"
}

func init() {
	db.RegisterModel(&Language{})
}
