package services

import (
	"errors"
	"main/internal/models"
)

func ValidateActor(actor models.Actor) error {
	if actor.FirstName == "" || actor.LastName == "" {
		return errors.New("first name and last name are required")
	}
	return nil
}
