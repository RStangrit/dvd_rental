package actor

import (
	"main/internal/film_actor"
	"main/pkg/db"
	"time"

	"gorm.io/gorm"
)

type Actor struct {
	ActorID    int                    `json:"actor_id" gorm:"type: integer; primaryKey;autoIncrement;not null"`
	FirstName  string                 `json:"first_name" gorm:"type: varchar(45);not null"`
	LastName   string                 `json:"last_name" gorm:"type: varchar(45);not null"`
	LastUpdate time.Time              `json:"last_update" gorm:"type: timestamp;not null; autoUpdateTime;default:now()"`
	FilmActors []film_actor.FilmActor `json:"-" gorm:"foreignKey:ActorID"`
	DeletedAt  gorm.DeletedAt         `json:"deleted_at"`
}

func (Actor) TableName() string {
	return "actor"
}

func init() {
	db.RegisterModel(&Actor{})
}
