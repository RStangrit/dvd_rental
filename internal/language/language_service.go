package language

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type LanguageService struct {
	repo *LanguageRepository
}

func NewLanguageService(repo *LanguageRepository) *LanguageService {
	return &LanguageService{repo: repo}
}

func (service *LanguageService) CreateLanguage(newLanguage *Language) error {
	err := service.ValidateLanguage(newLanguage)
	if err != nil {
		return err
	} else {
		return service.repo.InsertLanguage(newLanguage)
	}
}

func (service *LanguageService) CreateLanguages(newLanguages []*Language) ([]string, []Language, error) {
	var validationErrors []string
	var createdLanguages []Language

	for _, newLanguage := range newLanguages {
		if err := service.ValidateLanguage(newLanguage); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err := service.repo.InsertLanguage(newLanguage); err != nil {
			return validationErrors, createdLanguages, err
		}

		createdLanguages = append(createdLanguages, *newLanguage)
	}
	return validationErrors, createdLanguages, nil
}

func (service *LanguageService) ReadAllLanguages(pagination db.Pagination, filters map[string]any) ([]Language, int64, error) {
	languages, totalRecords, err := service.repo.SelectAllLanguages(pagination, filters)
	if err != nil {
		return nil, 0, err
	}
	return languages, totalRecords, nil
}

func (service *LanguageService) ReadOneLanguage(languageId int64) (*Language, error) {
	language, err := service.repo.SelectOneLanguage(languageId)
	if err != nil {
		return nil, err
	}
	if language == nil {
		return nil, fmt.Errorf("Language not found")
	}
	return language, nil
}

func (service *LanguageService) UpdateOneLanguage(language *Language) error {
	err := service.ValidateLanguage(language)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneLanguage(*language)
	}
}

func (service *LanguageService) DeleteOneLanguage(language *Language) error {
	return service.repo.DeleteOneLanguage(*language)
}

var ErrInvalidLanguageName = errors.New("language name is required and must be less than or equal to 20 characters")

func (service *LanguageService) ValidateLanguage(language *Language) error {
	if language.Name == "" || len(language.Name) > 20 {
		return ErrInvalidLanguageName
	}
	return nil
}
