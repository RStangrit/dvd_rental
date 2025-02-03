package store

import (
	"main/internal/staff"
	"main/pkg/db"
	"time"

	"gorm.io/gorm"
)

type Store struct {
	StoreID        int            `json:"store_id" gorm:"type:integer;primaryKey;autoIncrement"`
	ManagerStaffID int16          `json:"manager_staff_id" gorm:"type:int2;not null;foreignKey:StaffID"`
	AddressID      int16          `json:"address_id" gorm:"type:int2;not null;foreignKey:AddressID"`
	LastUpdate     time.Time      `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Staff          staff.Staff    `json:"-"`
}

func (Store) TableName() string {
	return "store"
}

func init() {
	db.RegisterModel(&Store{})
}
