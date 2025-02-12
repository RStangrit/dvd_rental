package actor

import (
	"errors"
	"fmt"
	"testing"
)

func Test_ValidateActor(t *testing.T) {
	var actor = Actor{
		ActorID:   0,
		FirstName: "",
		LastName:  "",
	}

	err := ValidateActor(&actor)
	if err == nil {
		t.Errorf("ValidateActor(&actor) validation failed")
	}
}

func Test_ValidateActorV2(t *testing.T) {

	tests := []struct {
		name        string
		firstName   string
		lastName    string
		expectedErr error
	}{
		{"positive", "John", "Doe", nil},
		{"last name is empty", "John", "", ErrMissingName},
		{"first name is empty", "", "Doe", ErrMissingName},
		{"both first and last names are empty", "", "", ErrMissingName},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actor := Actor{
				FirstName: test.firstName,
				LastName:  test.lastName,
			}
			err := ValidateActor(&actor)

			if !errors.Is(err, test.expectedErr) {
				assertion := errors.Is(err, test.expectedErr)
				fmt.Println(assertion)
				t.Errorf("ValidateActor test failed: expected %v, received %v", test.expectedErr, err)
			}
		})
	}
}
