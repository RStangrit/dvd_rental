package payment

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreatePayment(db *gorm.DB, newPayment *Payment) error {
	return db.Table("payment").Create(&newPayment).Error
}

func ReadAllPayments(db *gorm.DB, pagination db.Pagination) ([]Payment, int64, error) {
	var payments []Payment
	var totalRecords int64

	db.Table("payment").Count(&totalRecords)
	err := db.Table("payment").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("payment_id asc").Find(&payments).Error
	return payments, totalRecords, err
}

func ReadOnePayment(db *gorm.DB, paymentID int64) (*Payment, error) {
	var payment Payment
	err := db.Table("payment").First(&payment, paymentID).Error
	return &payment, err
}

func UpdateOnePayment(db *gorm.DB, payment Payment) error {
	return db.Table("payment").Where("payment_id = ?", payment.PaymentID).Updates(payment).Error
}

func DeleteOnePayment(db *gorm.DB, payment Payment) error {
	return db.Table("payment").Where("payment_id = ?", payment.PaymentID).Delete(&payment).Error
}
