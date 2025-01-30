package models

import "time"

type Country struct {
	CountryID  int       `json:"country_id" gorm:"type: integer;primaryKey;autoIncrement, not null"`
	Country    string    `json:"country" gorm:"type: varchar(50) not null"`
	LastUpdate time.Time `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime"`
}

func (Country) TableName() string {
	return "country"
}
