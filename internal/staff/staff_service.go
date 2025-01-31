package staff

import (
	"errors"
	"regexp"
)

func ValidateStaff(staff *Staff) error {
	if staff.FirstName == "" || len(staff.FirstName) > 45 {
		return errors.New("first name is required and must be less than 45 characters")
	}
	if staff.LastName == "" || len(staff.LastName) > 45 {
		return errors.New("last name is required and must be less than 45 characters")
	}
	if staff.Email == "" || len(staff.Email) > 50 || !isValidStaffEmail(staff.Email) {
		return errors.New("valid email is required and must be less than 50 characters")
	}
	if staff.AddressID <= 0 {
		return errors.New("address ID must be a positive number")
	}
	if staff.StoreID <= 0 {
		return errors.New("store ID must be a positive number")
	}
	if staff.Active != true && staff.Active != false {
		return errors.New("active status must be true or false")
	}
	if staff.Username == "" || len(staff.Username) > 16 {
		return errors.New("username is required and must be less than 16 characters")
	}
	if staff.Password == "" || len(staff.Password) > 40 {
		return errors.New("password is required and must be less than 40 characters")
	}
	if len(staff.Picture) == 0 {
		return errors.New("picture is required")
	}
	return nil
}

func isValidStaffEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
