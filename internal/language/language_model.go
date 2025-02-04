package language

import (
	"main/internal/film"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type Language struct {
	LanguageID int            `json:"language_id" gorm:"type: integer;primaryKey;autoIncrement"`
	Name       string         `json:"name" gorm:"type: bpchar(20);not null;index"`
	LastUpdate time.Time      `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime;default:now()"`
	Films      []film.Film    `json:"-" gorm:"foreignKey:LanguageID"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (Language) TableName() string {
	return "language"
}

func init() {
	models.RegisterModel(&Language{})
}
