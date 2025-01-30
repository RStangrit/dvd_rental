package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

// CreatePayment создает новый платеж в базе данных.
func CreatePayment(newPayment *models.Payment) error {
	return db.GORM.Table("payment").Create(&newPayment).Error
}

// ReadAllPayments читает все платежи с учетом пагинации.
func ReadAllPayments(pagination db.Pagination) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var totalRecords int64

	db.GORM.Table("payment").Count(&totalRecords)
	err := db.GORM.Table("payment").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("payment_id asc").Find(&payments).Error
	return payments, totalRecords, err
}

// ReadOnePayment читает один платеж по его ID.
func ReadOnePayment(paymentID int64) (*models.Payment, error) {
	var payment models.Payment
	err := db.GORM.Table("payment").First(&payment, paymentID).Error
	return &payment, err
}

// UpdateOnePayment обновляет информацию о платеже.
func UpdateOnePayment(payment models.Payment) error {
	return db.GORM.Table("payment").Where("payment_id = ?", payment.PaymentID).Updates(payment).Error
}

// DeleteOnePayment удаляет платеж по его ID.
func DeleteOnePayment(payment models.Payment) error {
	return db.GORM.Table("payment").Where("payment_id = ?", payment.PaymentID).Delete(&payment).Error
}
