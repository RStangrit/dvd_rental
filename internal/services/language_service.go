package services

import (
	"errors"
	"main/internal/models"
)

func ValidateLanguage(language *models.Language) error {
	if language.Name == "" {
		return errors.New("language name is required")
	}

	if len(language.Name) > 20 {
		return errors.New("language name must be less than or equal to 20 characters")
	}

	return nil
}
