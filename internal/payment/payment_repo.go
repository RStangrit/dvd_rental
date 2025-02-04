package payment

import (
	"main/pkg/db"
)

func CreatePayment(newPayment *Payment) error {
	return db.GORM.Table("payment").Create(&newPayment).Error
}

func ReadAllPayments(pagination db.Pagination) ([]Payment, int64, error) {
	var payments []Payment
	var totalRecords int64

	db.GORM.Table("payment").Count(&totalRecords)
	err := db.GORM.Table("payment").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("payment_id asc").Find(&payments).Error
	return payments, totalRecords, err
}

func ReadOnePayment(paymentID int64) (*Payment, error) {
	var payment Payment
	err := db.GORM.Table("payment").First(&payment, paymentID).Error
	return &payment, err
}

func UpdateOnePayment(payment Payment) error {
	return db.GORM.Table("payment").Where("payment_id = ?", payment.PaymentID).Updates(payment).Error
}

func DeleteOnePayment(payment Payment) error {
	return db.GORM.Table("payment").Where("payment_id = ?", payment.PaymentID).Delete(&payment).Error
}
