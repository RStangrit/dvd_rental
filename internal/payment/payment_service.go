package payment

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type PaymentService struct {
	repo *PaymentRepository
}

func NewPaymentService(repo *PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (service *PaymentService) CreatePayment(newPayment *Payment) error {
	if err := service.ValidatePayment(newPayment); err != nil {
		return err
	}
	return service.repo.InsertPayment(newPayment)
}

func (service *PaymentService) ReadAllPayments(pagination db.Pagination) ([]Payment, int64, error) {
	payments, totalRecords, err := service.repo.SelectAllPayments(pagination)
	if err != nil {
		return nil, 0, err
	}
	return payments, totalRecords, nil
}

func (service *PaymentService) ReadOnePayment(paymentId int64) (*Payment, error) {
	payment, err := service.repo.SelectOnePayment(paymentId)
	if err != nil {
		return nil, err
	}
	if payment == nil {
		return nil, fmt.Errorf("payment not found")
	}
	return payment, nil
}

func (service *PaymentService) UpdateOnePayment(payment *Payment) error {
	if err := service.ValidatePayment(payment); err != nil {
		return err
	}
	return service.repo.UpdateOnePayment(*payment)
}

func (service *PaymentService) DeleteOnePayment(payment *Payment) error {
	return service.repo.DeleteOnePayment(*payment)
}

func (service *PaymentService) ValidatePayment(payment *Payment) error {
	if payment.CustomerID <= 0 {
		return errors.New("customer_id is required and must be a positive number")
	}
	if payment.StaffID <= 0 {
		return errors.New("staff_id is required and must be a positive number")
	}
	if payment.RentalID <= 0 {
		return errors.New("rental_id is required and must be a positive number")
	}
	if payment.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if payment.PaymentDate.IsZero() {
		return errors.New("payment_date is required")
	}
	return nil
}
