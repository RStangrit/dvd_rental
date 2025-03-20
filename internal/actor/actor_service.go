package actor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"main/pkg/db"
	redisClient "main/pkg/redis"
	"time"
)

type ActorService struct {
	repo  *ActorRepository
	redis *redisClient.RedisClient
}

func NewActorService(repo *ActorRepository, redis *redisClient.RedisClient) *ActorService {
	return &ActorService{repo: repo, redis: redis}
}

func (service *ActorService) CreateActor(newActor *Actor) error {
	err := service.ValidateActor(newActor)
	if err != nil {
		return err
	}

	err = service.repo.InsertActor(newActor)
	if err != nil {
		return err
	}

	ctx := context.Background()
	cacheKey := fmt.Sprintf("actor:%d", newActor.ActorID)

	actorJSON, err := json.Marshal(newActor)
	if err != nil {
		return err
	}

	err = service.redis.SetKey(ctx, cacheKey, string(actorJSON), 10*time.Minute)
	if err != nil {
		return err
	}

	return nil
}

func (service *ActorService) CreateActors(newActors []*Actor) ([]string, []Actor, error) {
	var validationErrors []string
	var createdActors []Actor
	ctx := context.Background()

	for _, newActor := range newActors {
		if err := service.ValidateActor(newActor); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err := service.repo.InsertActor(newActor); err != nil {
			return validationErrors, createdActors, err
		}

		createdActors = append(createdActors, *newActor)

		cacheKey := fmt.Sprintf("actor:%d", newActor.ActorID)

		actorJSON, err := json.Marshal(newActor)
		if err != nil {
			return validationErrors, createdActors, err
		}

		err = service.redis.SetKey(ctx, cacheKey, string(actorJSON), 10*time.Minute)
		if err != nil {
			return validationErrors, createdActors, err
		}
	}

	return validationErrors, createdActors, nil
}

func (service *ActorService) ReadAllActors(pagination db.Pagination) ([]Actor, int64, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("actors:page:%d:limit:%d", pagination.Page, pagination.Limit)

	if cachedData, err := service.redis.GetKey(ctx, cacheKey); err == nil {
		var cachedActors []Actor
		if err := json.Unmarshal([]byte(cachedData), &cachedActors); err == nil {
			totalRecords, err := service.repo.CountActors()
			if err != nil {
				return nil, 0, err
			}
			return cachedActors, totalRecords, nil
		}
	}

	actors, totalRecords, err := service.repo.SelectAllActors(pagination)
	if err != nil {
		return nil, 0, err
	}

	actorJSON, err := json.Marshal(actors)
	if err == nil {
		_ = service.redis.SetKey(ctx, cacheKey, string(actorJSON), time.Hour)
	}

	return actors, totalRecords, nil
}

func (service *ActorService) ReadOneActor(actorId int64) (*Actor, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("actor:%d", actorId)

	if cachedData, err := service.redis.GetKey(ctx, cacheKey); err == nil {
		var cachedActor Actor
		if err := json.Unmarshal([]byte(cachedData), &cachedActor); err == nil {
			return &cachedActor, nil
		}
	}

	actor, err := service.repo.SelectOneActor(actorId)
	if err != nil {
		return nil, err
	}
	if actor == nil {
		return nil, fmt.Errorf("Actor not found")
	}

	actorJSON, _ := json.Marshal(actor)
	service.redis.SetKey(ctx, cacheKey, string(actorJSON), 10*time.Minute)

	return actor, nil
}

func (service *ActorService) ReadOneActorFilms(actorId int64) (*Actor, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("actor:%d", actorId)

	if cachedData, err := service.redis.GetKey(ctx, cacheKey); err == nil {
		var cachedActor Actor
		if err := json.Unmarshal([]byte(cachedData), &cachedActor); err == nil {
			return &cachedActor, nil
		}
	}

	actor, err := service.repo.SelectOneActorFilms(actorId)
	if err != nil {
		return nil, err
	}
	if actor == nil {
		return nil, fmt.Errorf("Actor not found")
	}

	actorJSON, _ := json.Marshal(actor)
	service.redis.SetKey(ctx, cacheKey, string(actorJSON), 10*time.Minute)

	return actor, nil
}

func (service *ActorService) UpdateOneActor(actor *Actor) error {
	err := service.ValidateActor(actor)
	if err != nil {
		return err
	}

	err = service.repo.UpdateOneActor(*actor)
	if err != nil {
		return err
	}

	ctx := context.Background()
	cacheKey := fmt.Sprintf("actor:%d", actor.ActorID)

	actorJSON, err := json.Marshal(actor)
	if err != nil {
		return err
	}

	return service.redis.SetKey(ctx, cacheKey, string(actorJSON), 10*time.Minute)
}

func (service *ActorService) DeleteOneActor(actor *Actor) error {
	return service.repo.DeleteOneActor(*actor)
}

var ErrMissingName = errors.New("first name and last name are required")

func (service *ActorService) ValidateActor(actor *Actor) error {
	if actor.FirstName == "" || actor.LastName == "" {
		return ErrMissingName
	}

	ctx := context.Background()
	cacheKey := fmt.Sprintf("actor:%d", actor.ActorID)
	service.redis.DeleteKey(ctx, cacheKey)

	return nil
}
