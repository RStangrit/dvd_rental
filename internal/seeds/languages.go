package seeds

import "main/internal/models"

func ReturnLangSeeds() []models.Language {
	Languages := []models.Language{
		{Name: "English"},
		{Name: "Italian"},
		{Name: "Japanese"},
		{Name: "Mandarin"},
		{Name: "French"},
		{Name: "German"},
	}
	return Languages
}
