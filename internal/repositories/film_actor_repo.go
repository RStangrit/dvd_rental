package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateFilmActor(newFilmActor *models.FilmActor) error {
	return db.GORM.Table("film_actor").Create(&newFilmActor).Error
}

func ReadAllFilmActors(pagination db.Pagination) ([]models.FilmActor, int64, error) {
	var filmActors []models.FilmActor
	var totalRecords int64

	db.GORM.Table("film_actor").Count(&totalRecords)
	err := db.GORM.Table("film_actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc, film_id asc").Find(&filmActors).Error
	return filmActors, totalRecords, err
}

func ReadOneFilmActor(actorID, filmID int64) (*models.FilmActor, error) {
	var filmActor models.FilmActor
	err := db.GORM.Table("film_actor").Where("actor_id = ? AND film_id = ?", actorID, filmID).First(&filmActor).Error
	return &filmActor, err
}

func UpdateOneFilmActor(filmActor models.FilmActor) error {
	return db.GORM.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Updates(filmActor).Error
}

func DeleteOneFilmActor(filmActor models.FilmActor) error {
	return db.GORM.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Delete(&filmActor).Error
}
