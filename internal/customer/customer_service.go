package customer

import (
	"errors"
	"fmt"
	"main/pkg/db"
	"regexp"
)

type CustomerService struct {
	repo *CustomerRepository
}

func NewCustomerService(repo *CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (service *CustomerService) CreateCustomer(newCustomer *Customer) error {
	if err := service.ValidateCustomer(newCustomer); err != nil {
		return err
	}
	return service.repo.InsertCustomer(newCustomer)
}

func (service *CustomerService) ReadAllCustomers(pagination db.Pagination) ([]Customer, int64, error) {
	customers, totalRecords, err := service.repo.SelectAllCustomers(pagination)
	if err != nil {
		return nil, 0, err
	}
	return customers, totalRecords, nil
}

func (service *CustomerService) ReadOneCustomer(customerId int64) (*Customer, error) {
	customer, err := service.repo.SelectOneCustomer(customerId)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("customer not found")
	}
	return customer, nil
}

func (service *CustomerService) UpdateOneCustomer(customer *Customer) error {
	if err := service.ValidateCustomer(customer); err != nil {
		return err
	}
	return service.repo.UpdateOneCustomer(*customer)
}

func (service *CustomerService) DeleteOneCustomer(customer *Customer) error {
	return service.repo.DeleteOneCustomer(*customer)
}

func (service *CustomerService) ValidateCustomer(customer *Customer) error {
	if customer.FirstName == "" || len(customer.FirstName) > 45 {
		return errors.New("first name is required and must be less than 45 characters")
	}
	if customer.LastName == "" || len(customer.LastName) > 45 {
		return errors.New("last name is required and must be less than 45 characters")
	}
	if customer.Email == "" || len(customer.Email) > 50 || !isValidCustomerEmail(customer.Email) {
		return errors.New("valid email is required and must be less than 50 characters")
	}
	if customer.AddressID <= 0 {
		return errors.New("address ID must be a positive number")
	}
	if customer.StoreID <= 0 {
		return errors.New("store ID must be a positive number")
	}
	if customer.Active != 0 && customer.Active != 1 {
		return errors.New("active status must be 0 or 1")
	}
	return nil
}

func isValidCustomerEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
