package models

import "time"

type City struct {
	CityID     int16     `json:"city_id" gorm:"type:serial4;primaryKey;autoIncrement"`
	City       string    `json:"city" gorm:"type:varchar(50);not null"`
	CountryID  int16     `json:"country_id" gorm:"type:int2;not null;foreignKey:CountryID"`
	LastUpdate time.Time `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	Addresses  []Address `json:"-" gorm:"foreignKey:CityID"`
}

func (City) TableName() string {
	return "city"
}
