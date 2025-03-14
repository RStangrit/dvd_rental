package film_category

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type FilmCategoryRepository struct {
	db *gorm.DB
}

func NewFilmCategoryRepository(db *gorm.DB) *FilmCategoryRepository {
	return &FilmCategoryRepository{db: db}
}

func (repo *FilmCategoryRepository) InsertFilmCategory(newFilmCategory *FilmCategory) error {
	return repo.db.Table("film_category").Create(&newFilmCategory).Error
}

func (repo *FilmCategoryRepository) SelectAllFilmCategories(pagination db.Pagination, sortParams string) ([]FilmCategory, int64, error) {
	var filmCategories []FilmCategory
	var totalRecords int64

	repo.db.Table("film_category").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("film_category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(sortParams).Find(&filmCategories).Error
	return filmCategories, totalRecords, err
}

func (repo *FilmCategoryRepository) SelectOneFilmCategory(filmID, categoryID int64) (*FilmCategory, error) {
	var filmCategory FilmCategory
	err := repo.db.Table("film_category").Where("film_id = ? AND category_id = ?", filmID, categoryID).First(&filmCategory).Error
	return &filmCategory, err
}

func (repo *FilmCategoryRepository) UpdateOneFilmCategory(filmID, categoryID int, updatedFilmCategory *FilmCategory) error {
	return repo.db.Table("film_category").Where("film_id = ? AND category_id = ?", filmID, categoryID).
		Updates(updatedFilmCategory).Error
}

func (repo *FilmCategoryRepository) DeleteOneFilmCategory(filmCategory FilmCategory) error {
	return repo.db.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).
		Delete(&filmCategory).Error
}
