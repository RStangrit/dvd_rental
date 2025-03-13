package film

import (
	"errors"
	"fmt"
	"main/pkg/db"
	"math"
	"time"
)

type FilmService struct {
	repo *FilmRepository
}

func NewFilmService(repo *FilmRepository) *FilmService {
	return &FilmService{repo: repo}
}

func (service *FilmService) CreateFilm(newFilm *Film) error {
	if err := service.ValidateFilm(newFilm); err != nil {
		return err
	}
	return service.repo.InsertFilm(newFilm)
}

func (service *FilmService) CreateFilms(newFilms []*Film) ([]string, []Film, error) {
	var validationErrors []string
	var createdFilms []Film

	for _, newFilm := range newFilms {
		if err := service.ValidateFilm(newFilm); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err := service.repo.InsertFilm(newFilm); err != nil {
			return validationErrors, createdFilms, err
		}

		createdFilms = append(createdFilms, *newFilm)
	}
	return validationErrors, createdFilms, nil
}

func (service *FilmService) ReadAllFilms(pagination db.Pagination, filters FilmFilter) ([]Film, int64, error) {
	films, totalRecords, err := service.repo.SelectAllFilms(pagination, filters)
	if err != nil {
		return nil, 0, err
	}
	return films, totalRecords, nil
}

func (service *FilmService) ReadOneFilm(filmId int64) (*Film, error) {
	film, err := service.repo.SelectOneFilm(filmId)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, fmt.Errorf("Film not found")
	}
	return film, nil
}

func (service *FilmService) ReadOneFilmActors(filmId int64) (*Film, error) {
	film, err := service.repo.SelectOneFilmActors(filmId)
	if err != nil {
		return nil, err
	}
	if film == nil {
		return nil, fmt.Errorf("Film not found")
	}
	return film, nil
}

func (service *FilmService) UpdateOneFilm(film *Film) error {
	if err := service.ValidateFilm(film); err != nil {
		return err
	}
	return service.repo.UpdateOneFilm(*film)
}

func (service *FilmService) DeleteOneFilm(film *Film) error {
	return service.repo.DeleteOneFilm(*film)
}

func (service *FilmService) DiscountOneFilm(filmId int64, discount float64) error {
	if err := service.ValidateDiscountPercentage(discount); err != nil {
		return err
	}
	return service.repo.DiscountOneFilm(filmId, discount)
}

var ErrInvalidTitle = errors.New("title is required and must be less than 255 characters")
var ErrInvalidReleaseYear = errors.New("release year is invalid")
var ErrInvalidLanguageID = errors.New("language ID must be a positive number")
var ErrInvalidRentalDuration = errors.New("rental duration must be a positive number")
var ErrInvalidRentalRate = errors.New("rental rate is invalid")
var ErrInvalidLength = errors.New("length must be a positive number")
var ErrInvalidReplacementCost = errors.New("replacement cost is invalid")
var ErrInvalidRating = errors.New("invalid rating")
var ErrInvalidFeature = errors.New("invalid special feature")
var ErrInvalidDiscount = errors.New("discount percentage must be a whole number and between 1 and 99")

func (service *FilmService) ValidateFilm(film *Film) error {
	if film.Title == "" || len(film.Title) > 255 {
		return ErrInvalidTitle
	}
	if film.ReleaseYear < 1900 || film.ReleaseYear > time.Now().Year()+5 {
		return ErrInvalidReleaseYear
	}
	if film.LanguageID <= 0 {
		return ErrInvalidLanguageID
	}
	if film.RentalDuration <= 0 {
		return ErrInvalidRentalDuration
	}
	if film.RentalRate <= 0 || film.RentalRate > 99.99 {
		return ErrInvalidRentalRate
	}
	if film.Length <= 0 {
		return ErrInvalidLength
	}
	if film.ReplacementCost <= 0 || film.ReplacementCost > 999.99 {
		return ErrInvalidReplacementCost
	}
	if !isValidRating(film.Rating) {
		return ErrInvalidRating
	}
	for _, feature := range film.SpecialFeatures {
		if !isValidFeature(feature) {
			return ErrInvalidFeature
		}
	}
	return nil
}

func isValidRating(rating string) bool {
	validRatings := map[string]bool{"G": true, "PG": true, "PG-13": true, "R": true, "NC-17": true}
	_, exists := validRatings[rating]
	return exists
}

func isValidFeature(feature string) bool {
	validFeatures := map[string]bool{"Trailers": true, "Commentaries": true, "Deleted Scenes": true, "Behind the Scenes": true}
	_, exists := validFeatures[feature]
	return exists
}

func (service *FilmService) ValidateDiscountPercentage(discount float64) error {
	if discount != math.Trunc(discount) {
		return ErrInvalidDiscount
	}
	d := int(discount)
	if d < 1 || d > 99 {
		return ErrInvalidDiscount
	}
	return nil
}
