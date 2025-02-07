package film_actor

import (
	"errors"
)

func ValidateFilmActor(filmActor *FilmActor) error {
	if filmActor.ActorID <= 0 {
		return errors.New("actor_id must be a positive integer")
	}

	if filmActor.FilmID <= 0 {
		return errors.New("film_id must be a positive integer")
	}

	return nil
}
