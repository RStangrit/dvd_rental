package repositories

import (
	"errors"
	"main/internal/models"
	"main/pkg/db"
)

func CreateFilm(newFilm *models.Film) error {
	result := db.GORM.Table("film").Create(&newFilm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadAllFilms(pagination db.Pagination) ([]models.Film, int64, error) {
	var films []models.Film
	var totalRecords int64

	db.GORM.Table("film").Count(&totalRecords)

	result := db.GORM.Table("film").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("film_id asc").Find(&films)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, 0, errors.New("films not found")
	}

	return films, totalRecords, nil
}

func ReadOneFilm(filmId int64) (*models.Film, error) {
	var film models.Film
	result := db.GORM.Table("film").First(&film, filmId)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("film not found")
	}

	return &film, nil
}

func UpdateOneFilm(film models.Film) error {
	result := db.GORM.Table("film").Omit("id").Updates(film)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteOneFilm(film models.Film) error {
	result := db.GORM.Delete(film)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
