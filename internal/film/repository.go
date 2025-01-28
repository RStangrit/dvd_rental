package film

import (
	"main/pkg/db"
)

func (newFilm *Film) createFilm() error {
	result := db.GORM.Table("film").Create(&newFilm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
