package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateFilm(newFilm *models.Film) error {
	return db.GORM.Table("film").Create(&newFilm).Error
}

func ReadAllFilms(pagination db.Pagination) ([]models.Film, int64, error) {
	var films []models.Film
	var totalRecords int64

	db.GORM.Table("film").Count(&totalRecords)
	err := db.GORM.Table("film").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("film_id asc").Find(&films).Error
	return films, totalRecords, err
}

func ReadOneFilm(filmId int64) (*models.Film, error) {
	var film models.Film
	err := db.GORM.Table("film").First(&film, filmId).Error
	return &film, err
}

func UpdateOneFilm(film models.Film) error {
	return db.GORM.Table("film").Omit("id").Updates(film).Error
}

func DeleteOneFilm(film models.Film) error {
	return db.GORM.Delete(&film).Error
}
