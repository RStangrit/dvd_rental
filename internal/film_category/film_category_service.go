package film_category

import (
	"errors"
)

func ValidateFilmCategory(filmCategory *FilmCategory) error {
	if filmCategory.FilmID <= 0 {
		return errors.New("film_id must be a positive integer")
	}

	if filmCategory.CategoryID <= 0 {
		return errors.New("category_id must be a positive integer")
	}

	return nil
}
