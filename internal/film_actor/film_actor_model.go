package film_actor

import (
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type FilmActor struct {
	ActorID    int            `json:"actor_id" gorm:"primaryKey;not null;foreignKey:ActorID;index"`
	FilmID     int            `json:"film_id" gorm:"primaryKey;not null;foreignKey:FilmID;index"`
	LastUpdate time.Time      `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime;default:now()"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (FilmActor) TableName() string {
	return "film_actor"
}

func init() {
	models.RegisterModel(&FilmActor{})
}
