package city

import (
	"errors"
)

func ValidateCity(city *City) error {
	if city.City == "" {
		return errors.New("city name is required")
	}

	if len(city.City) > 50 {
		return errors.New("city name must be less than or equal to 50 characters")
	}

	if city.CountryID <= 0 {
		return errors.New("country_id must be a positive integer")
	}

	return nil
}
