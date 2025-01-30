package repositories

import (
	"errors"
	"main/internal/models"
	"main/pkg/db"
)

func CreateLanguage(newLanguage *models.Language) error {
	result := db.GORM.Table("language").Create(&newLanguage)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadAllLanguages(pagination db.Pagination) ([]models.Language, int64, error) {
	var languages []models.Language
	var totalRecords int64

	db.GORM.Table("language").Count(&totalRecords)

	result := db.GORM.Table("language").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("language_id asc").Find(&languages)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, 0, errors.New("languages not found")
	}

	return languages, totalRecords, nil
}

func ReadOneLanguage(languageId int64) (*models.Language, error) {
	var language models.Language
	result := db.GORM.Table("language").First(&language, languageId)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("language not found")
	}

	return &language, nil
}

func UpdateOneLanguage(language models.Language) error {
	result := db.GORM.Table("language").Omit("language_id").Updates(language)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteOneLanguage(language models.Language) error {
	result := db.GORM.Delete(&language)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
