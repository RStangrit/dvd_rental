package services

import (
	"errors"
	"main/internal/models"
	"time"
)

func ValidateRental(rental *models.Rental) error {
	if rental.RentalID <= 0 {
		return errors.New("rental_id must be a positive integer")
	}

	if rental.RentalDate.IsZero() {
		return errors.New("rental_date must not be empty")
	}
	if rental.RentalDate.After(time.Now()) {
		return errors.New("rental_date cannot be in the future")
	}

	if rental.InventoryID <= 0 {
		return errors.New("inventory_id must be a positive integer")
	}

	if rental.CustomerID <= 0 {
		return errors.New("customer_id must be a positive integer")
	}

	if rental.ReturnDate.IsZero() {
		return errors.New("return_date must not be empty")
	}
	if rental.ReturnDate.Before(rental.RentalDate) {
		return errors.New("return_date cannot be before rental_date")
	}

	if rental.StaffID <= 0 {
		return errors.New("staff_id must be a positive integer")
	}

	if rental.LastUpdate.IsZero() {
		return errors.New("last_update must not be empty")
	}

	return nil
}
