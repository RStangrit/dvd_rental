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
	LastUpdate time.Time      `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime"`
	Films      []film.Film    `json:"films" gorm:"foreignKey:LanguageID"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

type LanguageListResponse struct {
	Data  []Language `json:"data"`
	Limit int        `json:"limit" example:"10"`
	Page  int        `json:"page" example:"1"`
	Total int        `json:"total" example:"307"`
}

type ErrorResponse struct {
	Errors []string `json:"errors" example:"language name is required and must be less than or equal to 20 characters"`
}

func (Language) TableName() string {
	return "language"
}

func (language *Language) LoadFilms(db *gorm.DB) error {
	return db.Model(language).Association("Films").Find(&language.Films)
}

func init() {
	models.RegisterModel(&Language{})
}
