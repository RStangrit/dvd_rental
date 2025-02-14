package customer

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateCustomer(db *gorm.DB, newCustomer *Customer) error {
	return db.Table("customer").Create(&newCustomer).Error
}

func ReadAllCustomers(db *gorm.DB, pagination db.Pagination) ([]Customer, int64, error) {
	var customers []Customer
	var totalRecords int64

	db.Table("customer").Count(&totalRecords)
	err := db.Table("customer").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("customer_id asc").Find(&customers).Error
	return customers, totalRecords, err
}

func ReadOneCustomer(db *gorm.DB, customerId int64) (*Customer, error) {
	var customer Customer
	err := db.Table("customer").First(&customer, customerId).Error
	return &customer, err
}

func UpdateOneCustomer(db *gorm.DB, customer Customer) error {
	return db.Table("customer").Omit("customer_id").Updates(customer).Error
}

func DeleteOneCustomer(db *gorm.DB, customer Customer) error {
	return db.Delete(&customer).Error
}
