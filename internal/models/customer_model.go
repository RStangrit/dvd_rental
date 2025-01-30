package models

import "time"

type Customer struct {
	CustomerID int16     `json:"customer_id" gorm:"type:serial4;primaryKey;autoIncrement"`
	StoreID    int16     `json:"store_id" gorm:"type:int2;not null"`
	FirstName  string    `json:"first_name" gorm:"type:varchar(45);not null"`
	LastName   string    `json:"last_name" gorm:"type:varchar(45);not null"`
	Email      string    `json:"email" gorm:"type:varchar(50);not null"`
	AddressID  int16     `json:"address_id" gorm:"type:int2;not null;foreignKey:AddressID"`
	Activebool bool      `json:"activebool" gorm:"type:boolean;not null;default:true"`
	CreateDate time.Time `json:"create_date" gorm:"type:date;not null;default:current_date"`
	LastUpdate time.Time `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	Active     int       `json:"active" gorm:"type:int4;not null"`
	Rentals    []Rental  `json:"-" gorm:"foreignKey:CustomerID"`
	Payments   []Payment `json:"-" gorm:"foreignKey:CustomerID"`
}

func (Customer) TableName() string {
	return "customer"
}
