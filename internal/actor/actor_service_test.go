package actor

// import (
// 	"errors"
// 	"testing"
// )

// func Test_ValidateActor(t *testing.T) {

// 	tests := []struct {
// 		name        string
// 		firstName   string
// 		lastName    string
// 		expectedErr error
// 	}{
// 		{"positive", "John", "Doe", nil},
// 		{"last name is empty", "John", "", ErrMissingName},
// 		{"first name is empty", "", "Doe", ErrMissingName},
// 		{"both first and last names are empty", "", "", ErrMissingName},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			actor := Actor{
// 				FirstName: test.firstName,
// 				LastName:  test.lastName,
// 			}
// 			err := ValidateActor(&actor)

// 			if !errors.Is(err, test.expectedErr) {
// 				t.Errorf("ValidateActor test failed: expected %v, received %v", test.expectedErr, err)
// 			}
// 		})
// 	}
// }
