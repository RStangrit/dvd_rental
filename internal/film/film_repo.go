package film

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateFilm(db *gorm.DB, newFilm *Film) error {
	return db.Table("film").Create(&newFilm).Error
}

func ReadAllFilms(db *gorm.DB, pagination db.Pagination, filters FilmFilter) ([]Film, int64, error) {
	var films []Film
	var totalRecords int64

	db.Table("film").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("film").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("film_id asc").Where(filters).Find(&films).Error
	return films, totalRecords, err
}

func ReadOneFilm(db *gorm.DB, filmId int64) (*Film, error) {
	var film Film
	err := db.Table("film").First(&film, filmId).Error
	return &film, err
}

func ReadOneFilmActors(db *gorm.DB, filmId int64) (Film, error) {
	var film Film
	err := db.Preload("FilmActors").First(&film, filmId).Error
	if err != nil {
		return Film{}, err
	}
	return film, err
}

func UpdateOneFilm(db *gorm.DB, film Film) error {
	return db.Table("film").Omit("id").Updates(film).Error
}

func DeleteOneFilm(db *gorm.DB, film Film) error {
	return db.Delete(&film).Error
}

func DiscountOneFilm(db *gorm.DB, film Film, discount float64) error {
	return db.Table("film").
		Where("film_id = ?", &film.FilmID).
		Updates(map[string]interface{}{
			"rental_rate": gorm.Expr("rental_rate * (1 - CAST(? AS FLOAT) / 100)", discount),
		}).Error
}
