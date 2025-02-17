package city

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateCity(db *gorm.DB, newCity *City) error {
	return db.Table("city").Create(&newCity).Error
}

func ReadAllCities(db *gorm.DB, pagination db.Pagination) ([]City, int64, error) {
	var cities []City
	var totalRecords int64

	db.Table("city").Count(&totalRecords)
	err := db.Table("city").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("city_id asc").Find(&cities).Error
	return cities, totalRecords, err
}

func ReadOneCity(db *gorm.DB, cityId int64) (*City, error) {
	var city City
	err := db.Table("city").First(&city, cityId).Error
	return &city, err
}

func UpdateOneCity(db *gorm.DB, city City) error {
	return db.Table("city").Omit("city_id").Updates(city).Error
}

func DeleteOneCity(db *gorm.DB, city City) error {
	return db.Delete(&city).Error
}
