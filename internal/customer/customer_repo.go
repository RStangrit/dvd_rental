package customer

import (
	"main/pkg/db"
)

func CreateCustomer(newCustomer *Customer) error {
	return db.GORM.Table("customer").Create(&newCustomer).Error
}

func ReadAllCustomers(pagination db.Pagination) ([]Customer, int64, error) {
	var customers []Customer
	var totalRecords int64

	db.GORM.Table("customer").Count(&totalRecords)
	err := db.GORM.Table("customer").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("customer_id asc").Find(&customers).Error
	return customers, totalRecords, err
}

func ReadOneCustomer(customerId int64) (*Customer, error) {
	var customer Customer
	err := db.GORM.Table("customer").First(&customer, customerId).Error
	return &customer, err
}

func UpdateOneCustomer(customer Customer) error {
	return db.GORM.Table("customer").Omit("customer_id").Updates(customer).Error
}

func DeleteOneCustomer(customer Customer) error {
	return db.GORM.Delete(&customer).Error
}
