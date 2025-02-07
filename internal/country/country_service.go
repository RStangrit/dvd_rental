package country

import (
	"errors"
)

func ValidateCountry(country *Country) error {
	if country.Country == "" {
		return errors.New("country name is required")
	}

	if len(country.Country) > 50 {
		return errors.New("country name must be less than or equal to 50 characters")
	}

	return nil
}
