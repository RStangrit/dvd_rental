package film

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateFilm(newFilm *Film) error {
	return db.GORM.Table("film").Create(&newFilm).Error
}

func ReadAllFilms(pagination db.Pagination, filters FilmFilter) ([]Film, int64, error) {
	var films []Film
	var totalRecords int64

	db.GORM.Table("film").Count(&totalRecords)
	err := db.GORM.Table("film").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("film_id asc").Where(filters).Find(&films).Error
	return films, totalRecords, err
}

func ReadOneFilm(filmId int64) (*Film, error) {
	var film Film
	err := db.GORM.Table("film").First(&film, filmId).Error
	return &film, err
}

func ReadOneFilmActors(filmId int64) (Film, error) {
	var film Film
	err := db.GORM.Preload("FilmActors").First(&film, filmId).Error
	if err != nil {
		return Film{}, err
	}
	return film, err
}

func UpdateOneFilm(film Film) error {
	return db.GORM.Table("film").Omit("id").Updates(film).Error
}

func DeleteOneFilm(film Film) error {
	return db.GORM.Delete(&film).Error
}

func DiscountOneFilm(film Film, discount float64) error {
	return db.GORM.Table("film").
		Where("film_id = ?", &film.FilmID).
		Updates(map[string]interface{}{
			"rental_rate": gorm.Expr("rental_rate * (1 - CAST(? AS FLOAT) / 100)", discount),
		}).Error
}
