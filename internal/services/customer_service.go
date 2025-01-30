package services

import (
	"errors"
	"main/internal/models"
	"regexp"
)

func ValidateCustomer(customer *models.Customer) error {
	if customer.FirstName == "" || len(customer.FirstName) > 45 {
		return errors.New("first name is required and must be less than 45 characters")
	}
	if customer.LastName == "" || len(customer.LastName) > 45 {
		return errors.New("last name is required and must be less than 45 characters")
	}
	if customer.Email == "" || len(customer.Email) > 50 || !isValidEmail(customer.Email) {
		return errors.New("valid email is required and must be less than 50 characters")
	}
	if customer.AddressID <= 0 {
		return errors.New("address ID must be a positive number")
	}
	if customer.StoreID <= 0 {
		return errors.New("store ID must be a positive number")
	}
	if customer.Active != 0 && customer.Active != 1 {
		return errors.New("active status must be 0 or 1")
	}
	return nil
}

func isValidEmail(email string) bool {
	// Simple email validation regex
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
