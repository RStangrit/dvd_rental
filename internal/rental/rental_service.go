package rental

import (
	"errors"
	"fmt"
	"main/pkg/db"
	"time"
)

type RentalService struct {
	repo *RentalRepository
}

func NewRentalService(repo *RentalRepository) *RentalService {
	return &RentalService{repo: repo}
}

func (service *RentalService) CreateOneRental(newRental *Rental) error {
	if err := service.ValidateRental(newRental); err != nil {
		return err
	}
	return service.repo.InsertRental(newRental)
}

func (service *RentalService) ReadAllRentals(pagination db.Pagination) ([]Rental, int64, error) {
	rentals, totalRecords, err := service.repo.SelectAllRentals(pagination)
	if err != nil {
		return nil, 0, err
	}
	return rentals, totalRecords, nil
}

func (service *RentalService) ReadOneRental(rentalID int64) (*Rental, error) {
	rental, err := service.repo.SelectOneRental(rentalID)
	if err != nil {
		return nil, err
	}
	if rental == nil {
		return nil, fmt.Errorf("rental not found")
	}
	return rental, nil
}

func (service *RentalService) UpdateOneRental(rental *Rental) error {
	if err := service.ValidateRental(rental); err != nil {
		return err
	}
	return service.repo.UpdateOneRental(*rental)
}

func (service *RentalService) DeleteOneRental(rental *Rental) error {
	return service.repo.DeleteOneRental(*rental)
}

func (service *RentalService) ValidateRental(rental *Rental) error {
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
	return nil
}
