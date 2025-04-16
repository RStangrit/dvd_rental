package development

import (
	"errors"
	"fmt"
	"main/internal/city"
	"main/internal/country"
	"main/internal/film"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type DevelopmentRepository struct {
	db      *gorm.DB
	country *country.Country
	city    *city.City
	film    *film.Film
}

func NewDevelopmentRepository(db *gorm.DB) *DevelopmentRepository {
	return &DevelopmentRepository{db: db}
}

func (repo *DevelopmentRepository) MakeTransaction(country *country.Country, city *city.City) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&country).Error; err != nil {
			return err
		}

		city.CountryID = int16(country.CountryID)

		if err := tx.Create(&city).Error; err != nil {
			return err
		}

		return nil
	})
}

func (repo *DevelopmentRepository) SelectAllFilmsForIndexing(batchSize int) ([]film.Film, error) {
	var allFilms []film.Film
	var totalRecords int64
	offset := 0

	if repo.db == nil {
		fmt.Println("repo.db is nil!")
		return nil, errors.New("repo.db is nil")
	}

	err := repo.db.Table("film").Where("deleted_at IS NULL").Count(&totalRecords).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(totalRecords)

	for offset < int(totalRecords) {
		var foundFilms []film.Film

		err := repo.db.Table("film").
			Offset(offset).
			Limit(batchSize).
			Order("film_id asc").
			Where("deleted_at IS NULL").
			Find(&foundFilms).Error
		if err != nil {
			return nil, err
		}

		if len(foundFilms) == 0 {
			break
		}

		allFilms = append(allFilms, foundFilms...)
		offset += batchSize
	}

	fmt.Printf("Fetched %d records so far\n", len(allFilms))
	return allFilms, nil
}
