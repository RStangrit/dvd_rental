package user

import (
	"errors"
	"fmt"
	"main/pkg/auth"
	"main/pkg/db"
	"net/mail"
	"regexp"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) CreateUser(newUser *User) error {
	if err := service.ValidateUser(newUser); err != nil {
		return err
	}
	return service.repo.InsertUser(newUser)
}

func (service *UserService) ReadAllUsers(pagination db.Pagination) ([]User, int64, error) {
	users, totalRecords, err := service.repo.SelectAllUsers(pagination)
	if err != nil {
		return nil, 0, err
	}
	return users, totalRecords, nil
}

func (service *UserService) ReadOneUserByEmail(email string) (*User, error) {
	user, err := service.repo.SelectOneUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (service *UserService) ReadOneUserById(userId int64) (*User, error) {
	user, err := service.repo.SelectOneUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (service *UserService) UpdateOneUser(user *User) error {
	if err := service.ValidateUser(user); err != nil {
		return err
	}
	return service.repo.UpdateOneUser(*user)
}

func (service *UserService) DeleteOneUser(user *User) error {
	return service.repo.DeleteOneUser(*user)
}

func (service *UserService) LoginUser(email, password string) (string, error) {
	user, err := service.ReadOneUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := auth.CompareHashAndPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.CreateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UserService) LogoutUser(token string) error {
	if token == "" {
		return fmt.Errorf("missing token")
	}

	err := auth.VerifyToken(token)
	if err != nil {
		return err
	}

	auth.AddToBlacklist(token)
	return nil
}

func (service *UserService) ValidateUser(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if user.Email == "" || len(user.Email) > 50 || !isValidUserEmail(user.Email) {
		return errors.New("valid email is required and must be less than 50 characters")
	}

	if user.Password == "" || len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if !strongPassword(user.Password) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

func isValidUserEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func strongPassword(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[#?!@$%^&*-]`).MatchString(password)
	hasLength := len(password) >= 8

	return hasUpper && hasLower && hasNumber && hasSpecial && hasLength
}
