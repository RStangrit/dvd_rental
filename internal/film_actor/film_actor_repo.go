package film_actor

import (
	"main/pkg/db"
)

func CreateFilmActor(newFilmActor *FilmActor) error {
	return db.GORM.Table("film_actor").Create(&newFilmActor).Error
}

func ReadAllFilmActors(pagination db.Pagination) ([]FilmActor, int64, error) {
	var filmActors []FilmActor
	var totalRecords int64

	db.GORM.Table("film_actor").Count(&totalRecords)
	err := db.GORM.Table("film_actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc, film_id asc").Find(&filmActors).Error
	return filmActors, totalRecords, err
}

func ReadOneFilmActor(actorID, filmID int64) (*FilmActor, error) {
	var filmActor FilmActor
	err := db.GORM.Table("film_actor").Where("actor_id = ? AND film_id = ?", actorID, filmID).First(&filmActor).Error
	return &filmActor, err
}

func UpdateOneFilmActor(filmActor FilmActor) error {
	return db.GORM.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Updates(filmActor).Error
}

func DeleteOneFilmActor(filmActor FilmActor) error {
	return db.GORM.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Delete(&filmActor).Error
}
