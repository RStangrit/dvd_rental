package country

import (
	"main/internal/city"
	"main/pkg/db"
	"time"
)

type Country struct {
	CountryID  int         `json:"country_id" gorm:"type: integer;primaryKey;autoIncrement, not null"`
	Country    string      `json:"country" gorm:"type: varchar(50) not null"`
	LastUpdate time.Time   `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime;default:now()"`
	Cities     []city.City `json:"-" gorm:"foreignKey:CountryID"`
}

func (Country) TableName() string {
	return "country"
}

func init() {
	db.RegisterModel(&Country{})
}
