package city

import (
	"main/pkg/db"
)

func CreateCity(newCity *City) error {
	return db.GORM.Table("city").Create(&newCity).Error
}

func ReadAllCities(pagination db.Pagination) ([]City, int64, error) {
	var cities []City
	var totalRecords int64

	db.GORM.Table("city").Count(&totalRecords)
	err := db.GORM.Table("city").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("city_id asc").Find(&cities).Error
	return cities, totalRecords, err
}

func ReadOneCity(cityId int64) (*City, error) {
	var city City
	err := db.GORM.Table("city").First(&city, cityId).Error
	return &city, err
}

func UpdateOneCity(city City) error {
	return db.GORM.Table("city").Omit("city_id").Updates(city).Error
}

func DeleteOneCity(city City) error {
	return db.GORM.Delete(&city).Error
}
