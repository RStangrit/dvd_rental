package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateAddress(newAddress *models.Address) error {
	return db.GORM.Table("address").Create(&newAddress).Error
}

func ReadAllAddresses(pagination db.Pagination) ([]models.Address, int64, error) {
	var addresses []models.Address
	var totalRecords int64

	db.GORM.Table("address").Count(&totalRecords)
	err := db.GORM.Table("address").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("address_id asc").Find(&addresses).Error
	return addresses, totalRecords, err
}

func ReadOneAddress(addressId int64) (*models.Address, error) {
	var address models.Address
	err := db.GORM.Table("address").First(&address, addressId).Error
	return &address, err
}

func UpdateOneAddress(address models.Address) error {
	return db.GORM.Table("address").Omit("address_id").Updates(address).Error
}

func DeleteOneAddress(address models.Address) error {
	return db.GORM.Delete(&address).Error
}
