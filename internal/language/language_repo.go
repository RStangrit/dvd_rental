package language

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type LanguageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{db: db}
}

func (repo *LanguageRepository) InsertLanguage(newLanguage *Language) error {
	return repo.db.Table("language").Create(&newLanguage).Error
}

func (repo *LanguageRepository) SelectAllLanguages(pagination db.Pagination, filters map[string]any) ([]Language, int64, error) {
	var languages []Language
	var totalRecords int64

	repo.db.Table("language").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("language").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("language_id asc").Where(filters).Find(&languages).Error
	return languages, totalRecords, err
}

func (repo *LanguageRepository) SelectOneLanguage(languageId int64) (*Language, error) {
	var language Language
	err := repo.db.Table("language").First(&language, languageId).Error
	return &language, err
}

func (repo *LanguageRepository) UpdateOneLanguage(language Language) error {
	return repo.db.Table("language").Omit("language_id").Updates(language).Error
}

func (repo *LanguageRepository) DeleteOneLanguage(language Language) error {
	return repo.db.Delete(&language).Error
}
