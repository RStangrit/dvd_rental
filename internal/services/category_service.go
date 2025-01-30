package services

import (
	"errors"
	"main/internal/models"
)

func ValidateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	if len(category.Name) > 25 {
		return errors.New("category name must be less than or equal to 25 characters")
	}

	return nil
}
