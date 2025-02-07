package payment

import (
	"errors"
)

func ValidatePayment(payment *Payment) error {
	if payment.CustomerID == 0 {
		return errors.New("customer_id is required")
	}
	if payment.StaffID == 0 {
		return errors.New("staff_id is required")
	}
	if payment.RentalID == 0 {
		return errors.New("rental_id is required")
	}
	if payment.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if payment.PaymentDate.IsZero() {
		return errors.New("payment_date is required")
	}
	return nil
}
