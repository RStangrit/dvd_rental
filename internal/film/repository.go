package film

import (
	"errors"
	"main/pkg/db"
)

func (newFilm *Film) createFilm() error {
	result := db.GORM.Table("film").Create(&newFilm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func readAllFilms(pagination db.Pagination) ([]Film, int64, error) {
	var films []Film
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
