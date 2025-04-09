package customer

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (repo *CustomerRepository) InsertCustomer(newCustomer *Customer) error {
	return repo.db.Table("customer").Create(&newCustomer).Error
}

func (repo *CustomerRepository) SelectAllCustomers(pagination db.Pagination) ([]Customer, int64, error) {
	var customers []Customer
	var totalRecords int64

	repo.db.Table("customer").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("customer").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("customer_id asc").
		Find(&customers).Error

	return customers, totalRecords, err
}

func (repo *CustomerRepository) SelectOneCustomer(customerId int64) (*Customer, error) {
	var customer Customer
	err := repo.db.Table("customer").First(&customer, customerId).Error
	return &customer, err
}

func (repo *CustomerRepository) UpdateOneCustomer(customer Customer) error {
	return repo.db.Table("customer").Omit("customer_id").Updates(customer).Error
}

func (repo *CustomerRepository) DeleteOneCustomer(customer Customer) error {
	return repo.db.Delete(&customer).Error
}
