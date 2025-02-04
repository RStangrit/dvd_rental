package inventory

import (
	"main/internal/models"
	"main/internal/rental"
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	InventoryID int            `json:"inventory_id" gorm:"type:integer;primaryKey;autoIncrement"`
	FilmID      int16          `json:"film_id" gorm:"type:int2;not null;foreignKey:FilmID"`
	StoreID     int16          `json:"store_id" gorm:"type:int2;not null"`
	LastUpdate  time.Time      `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Rental      rental.Rental  `json:"-" gorm:"foreignKey:InventoryID"`
}

func (Inventory) TableName() string {
	return "inventory"
}

func init() {
	models.RegisterModel(&Inventory{})
}
