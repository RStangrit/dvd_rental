package models

import "time"

type FilmCategory struct {
	FilmID     int16     `json:"film_id" gorm:"type:int2;not null;foreignKey:FilmID"`
	CategoryID int16     `json:"category_id" gorm:"type:int2;not null;foreignKey:CategoryID"`
	LastUpdate time.Time `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
}

func (FilmCategory) TableName() string {
	return "film_category"
}
