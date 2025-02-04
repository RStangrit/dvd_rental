package customer

import (
	"main/internal/models"
	"main/internal/payment"
	"main/internal/rental"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerID int               `json:"customer_id" gorm:"type:integer;primaryKey;autoIncrement"`
	StoreID    int16             `json:"store_id" gorm:"type:int2;not null"`
	FirstName  string            `json:"first_name" gorm:"type:varchar(45);not null"`
	LastName   string            `json:"last_name" gorm:"type:varchar(45);not null"`
	Email      string            `json:"email" gorm:"type:varchar(50);not null;uniqueIndex"`
	AddressID  int16             `json:"address_id" gorm:"type:int2;not null;foreignKey:AddressID"`
	Activebool bool              `json:"activebool" gorm:"type:boolean;not null;default:true"`
	CreateDate time.Time         `json:"create_date" gorm:"type:date;not null;default:current_date"`
	LastUpdate time.Time         `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt  gorm.DeletedAt    `json:"deleted_at"`
	Active     int               `json:"active" gorm:"type:int4;not null"`
	Rentals    []rental.Rental   `json:"-" gorm:"foreignKey:CustomerID"`
	Payments   []payment.Payment `json:"-" gorm:"foreignKey:CustomerID"`
}

func (Customer) TableName() string {
	return "customer"
}

func init() {
	models.RegisterModel(&Customer{})
}
