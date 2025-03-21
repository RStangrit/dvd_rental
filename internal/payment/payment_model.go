package payment

import (
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	PaymentID   int            `json:"payment_id" gorm:"type:integer;primaryKey;autoIncrement"`
	CustomerID  int16          `json:"customer_id" gorm:"type:int2;not null;foreignKey:CustomerID"`
	StaffID     int16          `json:"staff_id" gorm:"type:int2;not null;foreignKey:StaffID"`
	RentalID    int32          `json:"rental_id" gorm:"type:int4;not null;foreignKey:RentalID"`
	Amount      float64        `json:"amount" gorm:"type:numeric(5,2);not null"`
	PaymentDate time.Time      `json:"payment_date" gorm:"type:timestamp;not null"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func (Payment) TableName() string {
	return "payment"
}

func init() {
	models.RegisterModel(&Payment{})
}
