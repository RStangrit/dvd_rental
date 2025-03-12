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

func (repo *ActorRepository) SelectAllActors(db *gorm.DB, pagination db.Pagination) ([]Actor, int64, error) {
	var actors []Actor
	var totalRecords int64

	repo.db.Table("actor").Count(&totalRecords)
	err := repo.db.Table("actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc").Find(&actors).Error
	return actors, totalRecords, err
}

func (repo *ActorRepository) SelectOneActor(db *gorm.DB, actorId int64) (*Actor, error) {
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

func (repo *ActorRepository) UpdateOneActor(db *gorm.DB, actor Actor) error {
	return repo.db.Table("actor").Omit("actor_id").Updates(actor).Error
}

func (repo *ActorRepository) DeleteOneActor(db *gorm.DB, actor Actor) error {
	return repo.db.Delete(&actor).Error
}
