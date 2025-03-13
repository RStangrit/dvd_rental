package country

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type CountryService struct {
	repo *CountryRepository
}

func NewCountryService(repo *CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (service *CountryService) CreateCountry(newCountry *Country) error {
	if err := service.ValidateCountry(newCountry); err != nil {
		return err
	}
	return service.repo.InsertCountry(newCountry)
}

func (service *CountryService) ReadAllCountries(pagination db.Pagination) ([]Country, int64, error) {
	countries, totalRecords, err := service.repo.SelectAllCountries(pagination)
	if err != nil {
		return nil, 0, err
	}
	return countries, totalRecords, nil
}

func (service *CountryService) ReadOneCountry(countryId int64) (*Country, error) {
	country, err := service.repo.SelectOneCountry(countryId)
	if err != nil {
		return nil, err
	}
	if country == nil {
		return nil, fmt.Errorf("Country not found")
	}
	return country, nil
}

func (service *CountryService) UpdateOneCountry(country *Country) error {
	if err := service.ValidateCountry(country); err != nil {
		return err
	}
	return service.repo.UpdateOneCountry(*country)
}

func (service *CountryService) DeleteOneCountry(country *Country) error {
	return service.repo.DeleteOneCountry(*country)
}

func (service *CountryService) ValidateCountry(country *Country) error {
	if country.Country == "" {
		return errors.New("country name is required")
	}

	if len(country.Country) > 50 {
		return errors.New("country name must be less than or equal to 50 characters")
	}

	return nil
}
