package country

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (repo *CountryRepository) InsertCountry(newCountry *Country) error {
	return repo.db.Table("country").Create(&newCountry).Error
}

func (repo *CountryRepository) SelectAllCountries(pagination db.Pagination) ([]Country, int64, error) {
	var countries []Country
	var totalRecords int64

	repo.db.Table("country").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("country").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("country_id asc").Find(&countries).Error
	return countries, totalRecords, err
}

func (repo *CountryRepository) SelectOneCountry(countryId int64) (*Country, error) {
	var country Country
	err := repo.db.Table("country").First(&country, countryId).Error
	return &country, err
}

func (repo *CountryRepository) UpdateOneCountry(country Country) error {
	return repo.db.Table("country").Omit("country_id").Updates(country).Error
}

func (repo *CountryRepository) DeleteOneCountry(country Country) error {
	return repo.db.Delete(&country).Error
}
