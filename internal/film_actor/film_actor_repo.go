package film_actor

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type FilmActorRepository struct {
	db *gorm.DB
}

func NewFilmActorRepository(db *gorm.DB) *FilmActorRepository {
	return &FilmActorRepository{db: db}
}

func (repo *FilmActorRepository) InsertFilmActor(newFilmActor *FilmActor) error {
	return repo.db.Table("film_actor").Create(&newFilmActor).Error
}

func (repo *FilmActorRepository) SelectAllFilmActors(pagination db.Pagination) ([]FilmActor, int64, error) {
	var filmActors []FilmActor
	var totalRecords int64

	repo.db.Table("film_actor").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("film_actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc, film_id asc").Find(&filmActors).Error
	return filmActors, totalRecords, err
}

func (repo *FilmActorRepository) SelectOneFilmActor(actorID, filmID int64) (*FilmActor, error) {
	var filmActor FilmActor
	err := repo.db.Table("film_actor").Where("actor_id = ? AND film_id = ?", actorID, filmID).First(&filmActor).Error
	return &filmActor, err
}

func (repo *FilmActorRepository) UpdateOneFilmActor(actorID, filmID int, updatedFilmActor *FilmActor) error {
	return repo.db.Table("film_actor").Where("actor_id = ? AND film_id = ?", actorID, filmID).Update("film_id", &updatedFilmActor.FilmID).Error
}

func (repo *FilmActorRepository) DeleteOneFilmActor(filmActor FilmActor) error {
	return repo.db.Table("film_actor").Where("actor_id = ? AND film_id = ?", filmActor.ActorID, filmActor.FilmID).Delete(&filmActor).Error
}
