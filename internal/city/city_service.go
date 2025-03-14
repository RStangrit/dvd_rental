package city

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type CityService struct {
	repo *CityRepository
}

func NewCityService(repo *CityRepository) *CityService {
	return &CityService{repo: repo}
}

func (service *CityService) CreateCity(newCity *City) error {
	err := service.ValidateCity(newCity)
	if err != nil {
		return err
	} else {
		return service.repo.InsertCity(newCity)
	}
}

func (service *CityService) ReadAllCities(pagination db.Pagination) ([]City, int64, error) {
	cities, totalRecords, err := service.repo.SelectAllCities(pagination)
	if err != nil {
		return nil, 0, err
	}
	return cities, totalRecords, nil
}

func (service *CityService) ReadOneCity(cityId int64) (*City, error) {
	city, err := service.repo.SelectOneCity(cityId)
	if err != nil {
		return nil, err
	}
	if city == nil {
		return nil, fmt.Errorf("City not found")
	}
	return city, nil
}

func (service *CityService) UpdateOneCity(city *City) error {
	err := service.ValidateCity(city)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneCity(*city)
	}
}

func (service *CityService) DeleteOneCity(city *City) error {
	return service.repo.DeleteOneCity(*city)
}

func (service *CityService) ValidateCity(city *City) error {
	if city.City == "" {
		return errors.New("city name is required")
	}

	if len(city.City) > 50 {
		return errors.New("city name must be less than or equal to 50 characters")
	}

	if city.CountryID <= 0 {
		return errors.New("country_id must be a positive integer")
	}

	return nil
}
