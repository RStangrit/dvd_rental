package actor

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateActor(db *gorm.DB, newActor *Actor) error {
	return db.Table("actor").Create(&newActor).Error
}

func ReadAllActors(db *gorm.DB, pagination db.Pagination) ([]*Actor, int64, error) {
	var actors []*Actor
	var totalRecords int64

	db.Table("actor").Count(&totalRecords)
	err := db.Table("actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc").Find(&actors).Error
	return actors, totalRecords, err
}

func ReadOneActor(db *gorm.DB, actorId int64) (*Actor, error) {
	var actor Actor
	err := db.Table("actor").First(&actor, actorId).Error
	return &actor, err
}

func ReadOneActorFilms(db *gorm.DB, actorId int64) (*Actor, error) {
	var actor Actor
	err := db.Preload("ActorFilms").
		Where("actor.actor_id = ?", actorId).
		First(&actor).Error

	return &actor, err
}

func UpdateOneActor(db *gorm.DB, actor Actor) error {
	return db.Table("actor").Omit("actor_id").Updates(actor).Error
}

func DeleteOneActor(db *gorm.DB, actor Actor) error {
	return db.Delete(&actor).Error
}
