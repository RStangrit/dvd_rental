package actor

import (
	"errors"
)

var ErrMissingName = errors.New("first name and last name are required")

func ValidateActor(actor *Actor) error {
	if actor.FirstName == "" || actor.LastName == "" {
		return ErrMissingName
	}
	return nil
}
