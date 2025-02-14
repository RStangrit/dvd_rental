package address

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateAddress(db *gorm.DB, newAddress *Address) error {
	return db.Table("address").Create(&newAddress).Error
}

func ReadAllAddresses(db *gorm.DB, pagination db.Pagination) ([]Address, int64, error) {
	var addresses []Address
	var totalRecords int64

	db.Table("address").Count(&totalRecords)
	err := db.Table("address").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("address_id asc").Find(&addresses).Error
	return addresses, totalRecords, err
}

func ReadOneAddress(db *gorm.DB, addressId int64) (*Address, error) {
	var address Address
	err := db.Table("address").First(&address, addressId).Error
	return &address, err
}

func UpdateOneAddress(db *gorm.DB, address Address) error {
	return db.Table("address").Omit("address_id").Updates(address).Error
}

func DeleteOneAddress(db *gorm.DB, address Address) error {
	return db.Delete(&address).Error
}
