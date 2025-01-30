package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateActor(newActor *models.Actor) error {
	return db.GORM.Table("actor").Create(&newActor).Error
}

func ReadAllActors(pagination db.Pagination) ([]models.Actor, int64, error) {
	var actors []models.Actor
	var totalRecords int64

	db.GORM.Table("actor").Count(&totalRecords)
	err := db.GORM.Table("actor").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("actor_id asc").Find(&actors).Error
	return actors, totalRecords, err
}

func ReadOneActor(actorId int64) (*models.Actor, error) {
	var actor models.Actor
	err := db.GORM.Table("actor").First(&actor, actorId).Error
	return &actor, err
}

func UpdateOneActor(actor models.Actor) error {
	return db.GORM.Table("actor").Omit("actor_id").Updates(actor).Error
}

func DeleteOneActor(actor models.Actor) error {
	return db.GORM.Delete(&actor).Error
}
