package models

import "time"

type Actor struct {
	ActorID    int       `json:"actor_id" gorm:"type: serial4; primaryKey;autoIncrement, not null"`
	FirstName  string    `json:"first_name" gorm:"type: varchar(45) not null"`
	LastName   string    `json:"last_name" gorm:"type: varchar(45) not null"`
	LastUpdate time.Time `json:"last_update" gorm:"type: timestamp not null; autoUpdateTime"`
}

func (Actor) TableName() string {
	return "actor"
}
