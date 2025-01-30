package services

import (
	"errors"
	"main/internal/models"
)

func ValidateAddress(address *models.Address) error {
	if address.Address == "" {
		return errors.New("address is required")
	}

	if len(address.Address) > 50 {
		return errors.New("address must be less than or equal to 50 characters")
	}

	if address.District == "" {
		return errors.New("district is required")
	}

	if len(address.District) > 20 {
		return errors.New("district must be less than or equal to 20 characters")
	}

	if address.CityID <= 0 {
		return errors.New("city_id must be a positive integer")
	}

	if address.Phone == "" {
		return errors.New("phone number is required")
	}

	if len(address.Phone) > 20 {
		return errors.New("phone number must be less than or equal to 20 characters")
	}

	return nil
}
