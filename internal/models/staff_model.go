package models

import "time"

type Staff struct {
	StaffID    int16     `json:"staff_id" gorm:"type:serial4;primaryKey;autoIncrement"`
	FirstName  string    `json:"first_name" gorm:"type:varchar(45);not null"`
	LastName   string    `json:"last_name" gorm:"type:varchar(45);not null"`
	AddressID  int16     `json:"address_id" gorm:"type:int2;not null;foreignKey:AddressID"`
	Email      string    `json:"email" gorm:"type:varchar(50);not null"`
	StoreID    int16     `json:"store_id" gorm:"type:int2;not null"`
	Active     bool      `json:"active" gorm:"type:boolean;not null;default:true"`
	Username   string    `json:"username" gorm:"type:varchar(16);not null"`
	Password   string    `json:"password" gorm:"type:varchar(40);not null"`
	LastUpdate time.Time `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	Picture    []byte    `json:"picture" gorm:"type:bytea;not null"`
	Rentals    []Rental  `json:"-" gorm:"foreignKey:StaffID"`
	Payments   []Payment `json:"-" gorm:"foreignKey:StaffID"`
}

func (Staff) TableName() string {
	return "staff"
}
