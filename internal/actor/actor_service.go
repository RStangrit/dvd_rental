package actor

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type ActorService struct {
	repo *ActorRepository
}

func NewActorService(repo *ActorRepository) *ActorService {
	return &ActorService{repo: repo}
}

func (service *ActorService) CreateActor(newActor *Actor) error {
	err := service.ValidateActor(newActor)
	if err != nil {
		return err
	} else {
		return service.repo.InsertActor(newActor)
	}
}

func (service *ActorService) CreateActors(newActors []*Actor) ([]string, []Actor, error) {
	var validationErrors []string
	var createdActors []Actor

	for _, newActor := range newActors {
		if err := service.ValidateActor(newActor); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err := service.repo.InsertActor(newActor); err != nil {
			return validationErrors, createdActors, err
		}

		createdActors = append(createdActors, *newActor)
	}
	return validationErrors, createdActors, nil
}

func (service *ActorService) ReadAllActors(pagination db.Pagination) ([]Actor, int64, error) {
	actors, totalRecords, err := service.repo.SelectAllActors(service.repo.db, pagination)
	if err != nil {
		return nil, 0, err
	}
	return actors, totalRecords, nil
}

func (service *ActorService) ReadOneActor(actorId int64) (*Actor, error) {
	actor, err := service.repo.SelectOneActor(service.repo.db, actorId)
	if err != nil {
		return nil, err
	}
	if actor == nil {
		return nil, fmt.Errorf("Actor not found")
	}
	return actor, nil
}

func (service *ActorService) ReadOneActorFilms(actorId int64) (*Actor, error) {
	actor, err := service.repo.SelectOneActorFilms(actorId)
	if err != nil {
		return nil, err
	}
	if actor == nil {
		return nil, fmt.Errorf("Actor not found")
	}
	return actor, nil
}

func (service *ActorService) UpdateOneActor(actor *Actor) error {
	err := service.ValidateActor(actor)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneActor(service.repo.db, *actor)
	}
}

func (service *ActorService) DeleteOneActor(actor *Actor) error {
	return service.repo.DeleteOneActor(service.repo.db, *actor)
}

var ErrMissingName = errors.New("first name and last name are required")

func (service *ActorService) ValidateActor(actor *Actor) error {
	if actor.FirstName == "" || actor.LastName == "" {
		return ErrMissingName
	}
	return nil
}
