package film_category

import (
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type FilmCategory struct {
	FilmID     int16          `json:"film_id" gorm:"type:int2;not null;foreignKey:FilmID;index"`
	CategoryID int16          `json:"category_id" gorm:"type:int2;not null;foreignKey:CategoryID;index"`
	LastUpdate time.Time      `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (FilmCategory) TableName() string {
	return "film_category"
}

func init() {
	models.RegisterModel(&FilmCategory{})
}
