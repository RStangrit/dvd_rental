package actor

import (
	"main/pkg/db"
)

func CreateActor(newActor *Actor) error {
	return db.GORM.Table("actor").Create(&newActor).Error
}

func ReadAllActors(pagination db.Pagination) ([]Actor, int64, error) {
	var actors []Actor
	var totalRecords int64

	db.GORM.Table("actor").Count(&totalRecords)
	err := db.GORM.Table("actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc").Find(&actors).Error
	return actors, totalRecords, err
}

func ReadOneActor(actorId int64) (*Actor, error) {
	var actor Actor
	err := db.GORM.Table("actor").First(&actor, actorId).Error
	return &actor, err
}

func ReadOneActorFilms(actorId int64) (Actor, error) {
	var actor Actor
	err := db.GORM.Preload("ActorFilms").
		Where("actor.actor_id = ?", actorId).
		First(&actor).Error

	if err != nil {
		return Actor{}, err
	}
	return actor, err
}

func UpdateOneActor(actor Actor) error {
	return db.GORM.Table("actor").Omit("actor_id").Updates(actor).Error
}

func DeleteOneActor(actor Actor) error {
	return db.GORM.Delete(&actor).Error
}
