package category

import (
	"main/internal/film_category"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CategoryID   int                          `json:"category_id" gorm:"type:integer;primaryKey;autoIncrement"`
	Name         string                       `json:"name" gorm:"type:varchar(25);not null"`
	LastUpdate   time.Time                    `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt    gorm.DeletedAt               `json:"-"`
	FilmCategory []film_category.FilmCategory `json:"-" gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "category"
}

func init() {
	models.RegisterModel(&Category{})
}
