package address

import (
	"main/pkg/db"
)

func CreateAddress(newAddress *Address) error {
	return db.GORM.Table("address").Create(&newAddress).Error
}

func ReadAllAddresses(pagination db.Pagination) ([]Address, int64, error) {
	var addresses []Address
	var totalRecords int64

	db.GORM.Table("address").Count(&totalRecords)
	err := db.GORM.Table("address").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("address_id asc").Find(&addresses).Error
	return addresses, totalRecords, err
}

func ReadOneAddress(addressId int64) (*Address, error) {
	var address Address
	err := db.GORM.Table("address").First(&address, addressId).Error
	return &address, err
}

func UpdateOneAddress(address Address) error {
	return db.GORM.Table("address").Omit("address_id").Updates(address).Error
}

func DeleteOneAddress(address Address) error {
	return db.GORM.Delete(&address).Error
}
