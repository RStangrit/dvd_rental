package models

import "time"

type Inventory struct {
	InventoryID int       `json:"inventory_id" gorm:"type:serial4;primaryKey;autoIncrement"`
	FilmID      int16     `json:"film_id" gorm:"type:int2;not null;foreignKey:FilmID"`
	StoreID     int16     `json:"store_id" gorm:"type:int2;not null"`
	LastUpdate  time.Time `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
}

func (Inventory) TableName() string {
	return "inventory"
}
