package development

import (
	"errors"
	"main/internal/city"
	"main/internal/country"
	"math/rand"
)

type DevelopmentService struct {
	repo *DevelopmentRepository
}

func NewDevelopmentService(repo *DevelopmentRepository) *DevelopmentService {
	return &DevelopmentService{repo: repo}
}

func (service *DevelopmentService) CreateTransaction(countryName, cityName string) error {
	newCountry := &country.Country{Country: countryName}
	newCity := &city.City{City: cityName}

	if err := service.ValidateCountry(newCountry); err != nil {
		return err
	}

	return service.repo.MakeTransaction(newCountry, newCity)
}

func (service *DevelopmentService) ValidateCountry(country *country.Country) error {
	if country.Country == "" {
		return errors.New("country name is required")
	}

	if len(country.Country) > 50 {
		return errors.New("country name must be less than or equal to 50 characters")
	}

	return nil
}

func (service *DevelopmentService) GenerateRandomString(stringLength int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, stringLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
