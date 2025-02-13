package actor

import (
	"errors"
)

func ValidateActor(actor *Actor) error {
	if actor.FirstName == "" || actor.LastName == "" {
		return errors.New("first name and last name are required")
	}
	return nil
}
