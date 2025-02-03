package film

import (
	"main/pkg/db"
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

func UpdateOneFilm(film Film) error {
	return db.GORM.Table("film").Omit("id").Updates(film).Error
}

func DeleteOneFilm(film Film) error {
	return db.GORM.Delete(&film).Error
}
