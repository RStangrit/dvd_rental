package film_actor

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type FilmActorService struct {
	repo *FilmActorRepository
}

func NewFilmActorService(repo *FilmActorRepository) *FilmActorService {
	return &FilmActorService{repo: repo}
}

func (service *FilmActorService) CreateFilmActor(newFilmActor *FilmActor) error {
	err := service.ValidateFilmActor(newFilmActor)
	if err != nil {
		return err
	} else {
		return service.repo.InsertFilmActor(newFilmActor)
	}
}

func (service *FilmActorService) ReadAllFilmActors(pagination db.Pagination) ([]FilmActor, int64, error) {
	filmActors, totalRecords, err := service.repo.SelectAllFilmActors(pagination)
	if err != nil {
		return nil, 0, err
	}
	return filmActors, totalRecords, nil
}

func (service *FilmActorService) ReadOneFilmActor(filmId, actorId int64) (*FilmActor, error) {
	filmActor, err := service.repo.SelectOneFilmActor(filmId, actorId)
	if err != nil {
		return nil, err
	}
	if filmActor == nil {
		return nil, fmt.Errorf("filmActor not found")
	}
	return filmActor, nil
}

func (service *FilmActorService) UpdateOneFilmActor(actorID, filmID int, updatedFilmActor *FilmActor) error {
	err := service.ValidateFilmActor(updatedFilmActor)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneFilmActor(actorID, filmID, updatedFilmActor)
	}
}

func (service *FilmActorService) DeleteOneFilmActor(filmActor *FilmActor) error {
	return service.repo.DeleteOneFilmActor(*filmActor)
}

func (service *FilmActorService) ValidateFilmActor(filmActor *FilmActor) error {
	if filmActor.ActorID <= 0 {
		return errors.New("actor_id must be a positive integer")
	}

	if filmActor.FilmID <= 0 {
		return errors.New("film_id must be a positive integer")
	}

	return nil
}
