package country

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateCountry(db *gorm.DB, newCountry *Country) error {
	return db.Table("country").Create(&newCountry).Error
}

func ReadAllCountries(db *gorm.DB, pagination db.Pagination) ([]Country, int64, error) {
	var countries []Country
	var totalRecords int64

	db.Table("country").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("country").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("country_id asc").Find(&countries).Error
	return countries, totalRecords, err
}

func ReadOneCountry(db *gorm.DB, countryId int64) (*Country, error) {
	var country Country
	err := db.Table("country").First(&country, countryId).Error
	return &country, err
}

func UpdateOneCountry(db *gorm.DB, country Country) error {
	return db.Table("country").Omit("country_id").Updates(country).Error
}

func DeleteOneCountry(db *gorm.DB, country Country) error {
	return db.Delete(&country).Error
}
