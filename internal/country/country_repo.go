package country

import (
	"main/pkg/db"
)

func CreateCountry(newCountry *Country) error {
	return db.GORM.Table("country").Create(&newCountry).Error
}

func ReadAllCountries(pagination db.Pagination) ([]Country, int64, error) {
	var countries []Country
	var totalRecords int64

	db.GORM.Table("country").Count(&totalRecords)
	err := db.GORM.Table("country").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("country_id asc").Find(&countries).Error
	return countries, totalRecords, err
}

func ReadOneCountry(countryId int64) (*Country, error) {
	var country Country
	err := db.GORM.Table("country").First(&country, countryId).Error
	return &country, err
}

func UpdateOneCountry(country Country) error {
	return db.GORM.Table("country").Omit("country_id").Updates(country).Error
}

func DeleteOneCountry(country Country) error {
	return db.GORM.Delete(&country).Error
}
