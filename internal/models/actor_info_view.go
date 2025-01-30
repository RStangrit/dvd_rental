package models

type ActorInfo struct {
	ActorID   int    `json:"actor_id" gorm:"type:int;not null"`
	FirstName string `json:"first_name" gorm:"type:varchar(45);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(45);not null"`
	FilmInfo  string `json:"film_info" gorm:"type:text;not null"`
}

func (ActorInfo) TableName() string {
	return "actor_info"
}
