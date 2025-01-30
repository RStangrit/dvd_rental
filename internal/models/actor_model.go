package models

import "time"

type Actor struct {
	ActorID    int         `json:"actor_id" gorm:"type: integer; primaryKey;autoIncrement, not null"`
	FirstName  string      `json:"first_name" gorm:"type: varchar(45) not null"`
	LastName   string      `json:"last_name" gorm:"type: varchar(45) not null"`
	LastUpdate time.Time   `json:"last_update" gorm:"type: timestamp not null; autoUpdateTime"`
	FilmActors []FilmActor `json:"-" gorm:"foreignKey:ActorID"`
}

func (Actor) TableName() string {
	return "actor"
}
