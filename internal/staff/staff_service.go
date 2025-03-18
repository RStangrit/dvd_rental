package staff

import (
	"errors"
	"fmt"
	"main/pkg/avatars_generator"
	"main/pkg/db"
	"regexp"
)

type StaffService struct {
	repo *StaffRepository
}

func NewStaffService(repo *StaffRepository) *StaffService {
	return &StaffService{repo: repo}
}

func (service *StaffService) CreateStaff(newStaff *Staff) error {
	if err := service.ValidateStaff(newStaff); err != nil {
		return err
	}

	newStaff.Picture = service.GenerateAvatar(newStaff.FirstName, newStaff.LastName)

	return service.repo.InsertStaff(newStaff)
}

func (service *StaffService) ReadAllStaff(pagination db.Pagination) ([]Staff, int64, error) {
	staffs, totalRecords, err := service.repo.SelectAllStaffs(pagination)
	if err != nil {
		return nil, 0, err
	}
	return staffs, totalRecords, nil
}

func (service *StaffService) ReadOneStaff(staffId int64) (*Staff, error) {
	staff, err := service.repo.SelectOneStaff(staffId)
	if err != nil {
		return nil, err
	}
	if staff == nil {
		return nil, fmt.Errorf("staff not found")
	}
	return staff, nil
}

func (service *StaffService) UpdateOneStaff(staff *Staff) error {
	if err := service.ValidateStaff(staff); err != nil {
		return err
	}

	staff.Picture = service.GenerateAvatar(staff.FirstName, staff.LastName)

	return service.repo.UpdateOneStaff(*staff)
}

func (service *StaffService) DeleteOneStaff(staff *Staff) error {
	return service.repo.DeleteOneStaff(*staff)
}

func (service *StaffService) ValidateStaff(staff *Staff) error {
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
	if staff.Username == "" || len(staff.Username) > 16 {
		return errors.New("username is required and must be less than 16 characters")
	}
	if staff.Password == "" || len(staff.Password) > 40 {
		return errors.New("password is required and must be less than 40 characters")
	}
	return nil
}

func isValidStaffEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func (service *StaffService) GenerateAvatar(firstName, lastName string) []byte {
	firstNameLetter := firstLetter(firstName)
	lastNameLetter := firstLetter(lastName)
	image, _ := avatars_generator.CreateAvatar(firstNameLetter + lastNameLetter)
	return image
}

func firstLetter(s string) string {
	if s == "" {
		return ""
	}
	return string([]rune(s)[0])
}
