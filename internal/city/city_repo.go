package city

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (repo *CityRepository) InsertCity(newCity *City) error {
	return repo.db.Table("city").Create(&newCity).Error
}

func (repo *CityRepository) SelectAllCities(pagination db.Pagination) ([]City, int64, error) {
	var cities []City
	var totalRecords int64

	repo.db.Table("city").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("city").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("city_id asc").Find(&cities).Error
	return cities, totalRecords, err
}

func (repo *CityRepository) SelectOneCity(cityId int64) (*City, error) {
	var city City
	err := repo.db.Table("city").First(&city, cityId).Error
	return &city, err
}

func (repo *CityRepository) UpdateOneCity(city City) error {
	return repo.db.Table("city").Omit("city_id").Updates(city).Error
}

func (repo *CityRepository) DeleteOneCity(city City) error {
	return repo.db.Delete(&city).Error
}
