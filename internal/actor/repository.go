package actor

import (
	"errors"
	"main/pkg/db"
)

func (newActor *Actor) createActor() error {
	result := db.GORM.Table("actor").Create(&newActor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func readAllActors(pagination db.Pagination) ([]Actor, int64, error) {
	var actors []Actor
	var totalRecords int64

	db.GORM.Table("actor").Count(&totalRecords)

	result := db.GORM.Table("actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc").Find(&actors)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, 0, errors.New("actors not found")
	}

	return actors, totalRecords, nil
}

func readOneActor(actorId int64) (*Actor, error) {
	var actor Actor
	result := db.GORM.Table("actor").First(&actor, actorId)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("actor not found")
	}

	return &actor, nil
}

func (actor Actor) updateoneActor() error {
	result := db.GORM.Table("actor").Omit("id").Updates(actor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
