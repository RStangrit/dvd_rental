package development

import (
	"main/internal/city"
	"main/internal/country"

	"gorm.io/gorm"
)

type DevelopmentRepository struct {
	db      *gorm.DB
	country *country.Country
	city    *city.City
}

func NewDevelopmentRepository(db *gorm.DB) *DevelopmentRepository {
	return &DevelopmentRepository{db: db}
}

func (repo *DevelopmentRepository) MakeTransaction(country *country.Country, city *city.City) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&country).Error; err != nil {
			return err
		}

		city.CountryID = int16(country.CountryID)

		if err := tx.Create(&city).Error; err != nil {
			return err
		}

		return nil
	})
}
