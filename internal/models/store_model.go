package models

import "time"

type Store struct {
	StoreID        int16     `json:"store_id" gorm:"type:serial4;primaryKey;autoIncrement"`
	ManagerStaffID int16     `json:"manager_staff_id" gorm:"type:int2;not null"`
	AddressID      int16     `json:"address_id" gorm:"type:int2;not null;foreignKey:AddressID"`
	LastUpdate     time.Time `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
}

func (Store) TableName() string {
	return "store"
}
