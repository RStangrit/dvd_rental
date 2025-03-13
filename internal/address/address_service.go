package address

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type AddressService struct {
	repo *AddressRepository
}

func NewAddressService(repo *AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (service *AddressService) CreateAddress(newAddress *Address) error {
	err := service.ValidateAddress(newAddress)
	if err != nil {
		return err
	} else {
		return service.repo.InsertAddress(newAddress)
	}
}

func (service *AddressService) ReadAllAddresses(pagination db.Pagination) ([]Address, int64, error) {
	addresses, totalRecords, err := service.repo.SelectAllAddresses(pagination)
	if err != nil {
		return nil, 0, err
	}
	return addresses, totalRecords, nil
}

func (service *AddressService) ReadOneAddress(addressId int64) (*Address, error) {
	address, err := service.repo.SelectOneAddress(addressId)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, fmt.Errorf("address not found")
	}
	return address, nil
}

func (service *AddressService) UpdateOneAddress(address *Address) error {
	err := service.ValidateAddress(address)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneAddress(*address)
	}
}

func (service *AddressService) DeleteOneAddress(address *Address) error {
	return service.repo.DeleteOneAddress(*address)
}

func (service *AddressService) ValidateAddress(address *Address) error {
	if address.Address == "" {
		return errors.New("address is required")
	}

	if len(address.Address) > 50 {
		return errors.New("address must be less than or equal to 50 characters")
	}

	if address.District == "" {
		return errors.New("district is required")
	}

	if len(address.District) > 20 {
		return errors.New("district must be less than or equal to 20 characters")
	}

	if address.CityID <= 0 {
		return errors.New("city_id must be a positive integer")
	}

	if address.Phone == "" {
		return errors.New("phone number is required")
	}

	if len(address.Phone) > 20 {
		return errors.New("phone number must be less than or equal to 20 characters")
	}

	return nil
}
