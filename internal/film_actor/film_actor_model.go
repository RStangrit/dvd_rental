package film_actor

import (
	"main/pkg/db"
	"time"
)

type FilmActor struct {
	ActorID    int       `json:"actor_id" gorm:"primaryKey;not null;foreignKey:ActorID"`
	FilmID     int       `json:"film_id" gorm:"primaryKey;not null;foreignKey:FilmID"`
	LastUpdate time.Time `json:"last_update" gorm:"type:timestamp;not null;autoUpdateTime"`
}

func (FilmActor) TableName() string {
	return "film_actor"
}

func init() {
	db.RegisterModel(&FilmActor{})
}
