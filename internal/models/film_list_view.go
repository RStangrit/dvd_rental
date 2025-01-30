package models

type FilmList struct {
	FID         int     `json:"fid" gorm:"type:int;not null"`
	Title       string  `json:"title" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text;not null"`
	Category    string  `json:"category" gorm:"type:varchar(25);not null"`
	Price       float64 `json:"price" gorm:"type:numeric(4,2);not null"`
	Length      int16   `json:"length" gorm:"type:int2;not null"`
	Rating      string  `json:"rating" gorm:"type:mpaa_rating;not null"`
	Actors      string  `json:"actors" gorm:"type:text;not null"`
}

func (FilmList) TableName() string {
	return "film_list"
}
