package actor

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type ActorRepository struct {
	db *gorm.DB
}

func NewActorRepository(db *gorm.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (repo *ActorRepository) InsertActor(newActor *Actor) error {
	return repo.db.Table("actor").Create(&newActor).Error
}

func (repo *ActorRepository) SelectAllActors(pagination db.Pagination) ([]Actor, int64, error) {
	var actors []Actor
	var totalRecords int64

	repo.db.Table("actor").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc").Find(&actors).Error
	return actors, totalRecords, err
}

func (repo *ActorRepository) SelectOneActor(actorId int64) (*Actor, error) {
	var actor Actor
	err := repo.db.Table("actor").First(&actor, actorId).Error
	return &actor, err
}

func (repo *ActorRepository) SelectOneActorFilms(actorId int64) (*Actor, error) {
	var actor Actor
	err := repo.db.Preload("ActorFilms").
		Where("actor.actor_id = ?", actorId).
		First(&actor).Error

	return &actor, err
}

func (repo *ActorRepository) UpdateOneActor(actor Actor) error {
	return repo.db.Table("actor").Omit("actor_id").Updates(actor).Error
}

func (repo *ActorRepository) DeleteOneActor(actor Actor) error {
	return repo.db.Delete(&actor).Error
}

func (repo *ActorRepository) CountActors() (int64, error) {
	var count int64
	err := repo.db.Table("actor").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
