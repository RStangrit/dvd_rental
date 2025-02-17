package address

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (repo *AddressRepository) InsertAddress(newAddress *Address) error {
	return repo.db.Table("address").Create(&newAddress).Error
}

func (repo *AddressRepository) SelectAllAddresses(db *gorm.DB, pagination db.Pagination) ([]Address, int64, error) {
	var addresses []Address
	var totalRecords int64

	db.Table("address").Count(&totalRecords)
	err := db.Table("address").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("address_id asc").Find(&addresses).Error
	return addresses, totalRecords, err
}

func (repo *AddressRepository) SelectOneAddress(db *gorm.DB, addressId int64) (*Address, error) {
	var address Address
	err := db.Table("address").First(&address, addressId).Error
	return &address, err
}

func (repo *AddressRepository) UpdateOneAddress(db *gorm.DB, address Address) error {
	return db.Table("address").Omit("address_id").Updates(address).Error
}

func (repo *AddressRepository) DeleteOneAddress(db *gorm.DB, address Address) error {
	return db.Delete(&address).Error
}
