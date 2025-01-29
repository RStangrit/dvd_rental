package repositories

import (
	"errors"
	"main/internal/models"
	"main/pkg/db"
)

func CreateActor(newActor *models.Actor) error {
	result := db.GORM.Table("actor").Create(&newActor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadAllActors(pagination db.Pagination) ([]models.Actor, int64, error) {
	var actors []models.Actor
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

func ReadOneActor(actorId int64) (*models.Actor, error) {
	var actor models.Actor
	result := db.GORM.Table("actor").First(&actor, actorId)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("actor not found")
	}

	return &actor, nil
}

func UpdateOneActor(actor models.Actor) error {
	result := db.GORM.Table("actor").Omit("id").Updates(actor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteOneActor(actor models.Actor) error {
	result := db.GORM.Delete(actor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
