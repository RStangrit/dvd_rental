package film

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type FilmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepository {
	return &FilmRepository{db: db}
}

func (repo *FilmRepository) InsertFilm(newFilm *Film) error {
	return repo.db.Table("film").Create(&newFilm).Error
}

func (repo *FilmRepository) SelectAllFilms(pagination db.Pagination, filters FilmFilter) ([]Film, int64, error) {
	var films []Film
	var totalRecords int64

	repo.db.Table("film").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("film").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("film_id asc").
		Where(filters).
		Find(&films).Error

	return films, totalRecords, err
}

func (repo *FilmRepository) SelectOneFilm(filmId int64) (*Film, error) {
	var film Film
	err := repo.db.Table("film").First(&film, filmId).Error
	return &film, err
}

func (repo *FilmRepository) SelectOneFilmActors(filmId int64) (*Film, error) {
	var film Film
	err := repo.db.Preload("FilmActors").
		Where("film.film_id = ?", filmId).
		First(&film).Error

	return &film, err
}

func (repo *FilmRepository) UpdateOneFilm(film Film) error {
	return repo.db.Table("film").Omit("film_id").Updates(film).Error
}

func (repo *FilmRepository) DeleteOneFilm(film Film) error {
	return repo.db.Delete(&film).Error
}

func (repo *FilmRepository) DiscountOneFilm(filmId int64, discount float64) error {
	return repo.db.Table("film").
		Where("film_id = ?", filmId).
		Updates(map[string]interface{}{
			"rental_rate": gorm.Expr("rental_rate * (1 - CAST(? AS FLOAT) / 100)", discount),
		}).Error
}
