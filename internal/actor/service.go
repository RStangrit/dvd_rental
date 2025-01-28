package actor

import "errors"

func (a *Actor) Validate() error {
	if a.FirstName == "" || a.LastName == "" {
		return errors.New("first name and last name are required")
	}
	return nil
}
