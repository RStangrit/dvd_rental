package language

import (
	"main/pkg/db"
)

func CreateLanguage(newLanguage *Language) error {
	return db.GORM.Table("language").Create(&newLanguage).Error
}

func ReadAllLanguages(pagination db.Pagination) ([]Language, int64, error) {
	var languages []Language
	var totalRecords int64

	db.GORM.Table("language").Count(&totalRecords)
	err := db.GORM.Table("language").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("language_id asc").Find(&languages).Error
	return languages, totalRecords, err
}

func ReadOneLanguage(languageId int64) (*Language, error) {
	var language Language
	err := db.GORM.Table("language").First(&language, languageId).Error
	return &language, err
}

func UpdateOneLanguage(language Language) error {
	return db.GORM.Table("language").Omit("language_id").Updates(language).Error
}

func DeleteOneLanguage(language Language) error {
	return db.GORM.Delete(&language).Error
}
