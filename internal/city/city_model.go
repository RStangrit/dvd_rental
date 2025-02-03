package city

import (
	"main/internal/address"
	"main/pkg/db"
	"time"

	"gorm.io/gorm"
)

type City struct {
	CityID     int               `json:"city_id" gorm:"type:integer;primaryKey;autoIncrement"`
	City       string            `json:"city" gorm:"type:varchar(50);not null"`
	CountryID  int16             `json:"country_id" gorm:"type:int2;not null;foreignKey:CountryID"`
	LastUpdate time.Time         `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt  gorm.DeletedAt    `json:"deleted_at"`
	Addresses  []address.Address `json:"-" gorm:"foreignKey:CityID"`
}

func (City) TableName() string {
	return "city"
}

func init() {
	db.RegisterModel(&City{})
}
