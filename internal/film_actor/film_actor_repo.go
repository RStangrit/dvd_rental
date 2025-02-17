package film_actor

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateFilmActor(db *gorm.DB, newFilmActor *FilmActor) error {
	return db.Table("film_actor").Create(&newFilmActor).Error
}

func ReadAllFilmActors(db *gorm.DB, pagination db.Pagination) ([]FilmActor, int64, error) {
	var filmActors []FilmActor
	var totalRecords int64

	db.Table("film_actor").Count(&totalRecords)
	err := db.Table("film_actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc, film_id asc").Find(&filmActors).Error
	return filmActors, totalRecords, err
}

func ReadOneFilmActor(db *gorm.DB, actorID, filmID int64) (*FilmActor, error) {
	var filmActor FilmActor
	err := db.Table("film_actor").Where("actor_id = ? AND film_id = ?", actorID, filmID).First(&filmActor).Error
	return &filmActor, err
}

func UpdateOneFilmActor(db *gorm.DB, filmActor FilmActor) error {
	return db.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Updates(filmActor).Error
}

func DeleteOneFilmActor(db *gorm.DB, filmActor FilmActor) error {
	return db.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Delete(&filmActor).Error
}
