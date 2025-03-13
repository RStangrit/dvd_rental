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

func (repo *AddressRepository) SelectAllAddresses(pagination db.Pagination) ([]Address, int64, error) {
	var addresses []Address
	var totalRecords int64

	repo.db.Table("address").Count(&totalRecords)
	err := repo.db.Table("address").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("address_id asc").Find(&addresses).Error
	return addresses, totalRecords, err
}

func (repo *AddressRepository) SelectOneAddress(addressId int64) (*Address, error) {
	var address Address
	err := repo.db.Table("address").First(&address, addressId).Error
	return &address, err
}

func (repo *AddressRepository) UpdateOneAddress(address Address) error {
	return repo.db.Table("address").Omit("address_id").Updates(address).Error
}

func (repo *AddressRepository) DeleteOneAddress(address Address) error {
	return repo.db.Delete(&address).Error
}
