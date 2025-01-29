package models

type Actor struct {
	ActorID   int    `json:"actor_id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (Actor) TableName() string {
	return "actor"
}
