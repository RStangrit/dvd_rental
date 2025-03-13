package actor

import (
	"main/internal/film_actor"
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type Actor struct {
	ActorID    int                    `json:"actor_id" gorm:"type: integer; primaryKey;autoIncrement;not null"`
	FirstName  string                 `json:"first_name" gorm:"type: varchar(45);not null"`
	LastName   string                 `json:"last_name" gorm:"type: varchar(45);not null"`
	LastUpdate time.Time              `json:"last_update" gorm:"type: timestamp;not null; autoUpdateTime;default:now()"`
	ActorFilms []film_actor.FilmActor `json:"actor_films" gorm:"foreignKey:ActorID"`
	DeletedAt  gorm.DeletedAt         `json:"-"`
}

var tableName = "actor"

func (Actor) TableName() string {
	return tableName
}

func registerModels() {
	models.RegisterModel(&Actor{})
}

func init() {
	registerModels()
}
