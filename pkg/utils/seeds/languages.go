package seeds

import (
	"main/internal/language"
	"main/pkg/db"
)

func langSeeds() []language.Language {
	Languages := []language.Language{
		{Name: "English"},
		{Name: "Italian"},
		{Name: "Japanese"},
		{Name: "Mandarin"},
		{Name: "French"},
		{Name: "German"},
	}
	return Languages
}

func SeedLanguageData() error {
	languages := langSeeds()
	if err := db.GORM.Create(&languages).Error; err != nil {
		return err
	}
	return nil
}
