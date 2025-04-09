package payment

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (repo *PaymentRepository) InsertPayment(newPayment *Payment) error {
	return repo.db.Table("payment").Create(&newPayment).Error
}

func (repo *PaymentRepository) SelectAllPayments(pagination db.Pagination) ([]Payment, int64, error) {
	var payments []Payment
	var totalRecords int64

	repo.db.Table("payment").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("payment").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("payment_id asc").
		Find(&payments).Error

	return payments, totalRecords, err
}

func (repo *PaymentRepository) SelectOnePayment(paymentID int64) (*Payment, error) {
	var payment Payment
	err := repo.db.Table("payment").First(&payment, paymentID).Error
	return &payment, err
}

func (repo *PaymentRepository) UpdateOnePayment(payment Payment) error {
	return repo.db.Table("payment").Where("payment_id = ?", payment.PaymentID).Updates(payment).Error
}

func (repo *PaymentRepository) DeleteOnePayment(payment Payment) error {
	return repo.db.Table("payment").Where("payment_id = ?", payment.PaymentID).Delete(&payment).Error
}
