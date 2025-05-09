package address

import (
	"main/internal/customer"
	"main/internal/models"
	"main/internal/staff"
	"main/internal/store"
	"time"

	"gorm.io/gorm"
)

type Address struct {
	AddressID  int               `json:"address_id" gorm:"type:integer;primaryKey;autoIncrement"`
	Address    string            `json:"address" gorm:"type:varchar(50);not null"`
	Address2   string            `json:"address2" gorm:"type:varchar(50)"`
	District   string            `json:"district" gorm:"type:varchar(20);not null"`
	CityID     int16             `json:"city_id" gorm:"type:int2;not null;foreignKey:CityID"`
	PostalCode string            `json:"postal_code" gorm:"type:varchar(10)"`
	Phone      string            `json:"phone" gorm:"type:varchar(20);not null"`
	LastUpdate time.Time         `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt    `json:"-"`
	Customer   customer.Customer `json:"-" gorm:"foreignKey:AddressID"`
	Staff      staff.Staff       `json:"-" gorm:"foreignKey:AddressID"`
	Store      store.Store       `json:"-" gorm:"foreignKey:AddressID"`
}

func (Address) TableName() string {
	return "address"
}

func init() {
	models.RegisterModel(&Address{})
}
