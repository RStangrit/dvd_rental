package rental

import (
	"main/internal/models"
	"main/internal/payment"
	"time"

	"gorm.io/gorm"
)

type Rental struct {
	RentalID    int             `json:"rental_id" gorm:"type:integer;primaryKey;autoIncrement"`
	RentalDate  time.Time       `json:"rental_date" gorm:"type:timestamp;not null"`
	InventoryID int32           `json:"inventory_id" gorm:"type:int4;not null;foreignKey:InventoryID"`
	CustomerID  int16           `json:"customer_id" gorm:"type:int2;not null;foreignKey:CustomerID"`
	ReturnDate  time.Time       `json:"return_date" gorm:"type:timestamp;not null"`
	StaffID     int16           `json:"staff_id" gorm:"type:int2;not null;foreignKey:StaffID"`
	LastUpdate  time.Time       `json:"last_update" gorm:"type:timestamp;not null;default:now()"`
	DeletedAt   gorm.DeletedAt  `json:"-"`
	Payment     payment.Payment `json:"-" gorm:"foreignKey:RentalID"`
}

func (Rental) TableName() string {
	return "rental"
}

func init() {
	models.RegisterModel(&Rental{})
}
