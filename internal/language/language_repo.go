package language

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateLanguage(db *gorm.DB, newLanguage *Language) error {
	return db.Table("language").Create(&newLanguage).Error
}

func ReadAllLanguages(db *gorm.DB, pagination db.Pagination, filters map[string]any) ([]Language, int64, error) {
	var languages []Language
	var totalRecords int64

	db.Table("language").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("language").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("language_id asc").Where(filters).Find(&languages).Error
	return languages, totalRecords, err
}

func ReadOneLanguage(db *gorm.DB, languageId int64) (*Language, error) {
	var language Language
	err := db.Table("language").First(&language, languageId).Error
	return &language, err
}

func UpdateOneLanguage(db *gorm.DB, language Language) error {
	return db.Table("language").Omit("language_id").Updates(language).Error
}

func DeleteOneLanguage(db *gorm.DB, language Language) error {
	return db.Delete(&language).Error
}
