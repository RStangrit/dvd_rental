package user

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

func ValidateUser(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	// Validate password strength (example)
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !strongPassword(user.Password) {
		fmt.Println(user.Password, !strongPassword(user.Password))
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number and one special character")
	}

	return nil
}

func strongPassword(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[#?!@$%^&*-]`).MatchString(password)
	hasLength := len(password) >= 8

	return hasUpper && hasLower && hasNumber && hasSpecial && hasLength
}
