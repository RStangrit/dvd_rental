package category

import (
	"errors"
)

func ValidateCategory(category *Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	if len(category.Name) > 25 {
		return errors.New("category name must be less than or equal to 25 characters")
	}

	return nil
}
