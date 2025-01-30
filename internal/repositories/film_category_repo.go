package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateFilmCategory(newFilmCategory *models.FilmCategory) error {
	return db.GORM.Table("film_category").Create(&newFilmCategory).Error
}

func ReadAllFilmCategories(pagination db.Pagination) ([]models.FilmCategory, int64, error) {
	var filmCategories []models.FilmCategory
	var totalRecords int64

	db.GORM.Table("film_category").Count(&totalRecords)
	err := db.GORM.Table("film_category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("film_id asc, category_id asc").Find(&filmCategories).Error
	return filmCategories, totalRecords, err
}

func ReadOneFilmCategory(filmID, categoryID int64) (*models.FilmCategory, error) {
	var filmCategory models.FilmCategory
	err := db.GORM.Table("film_category").Where("film_id = ? AND category_id = ?", filmID, categoryID).First(&filmCategory).Error
	return &filmCategory, err
}

func UpdateOneFilmCategory(filmCategory models.FilmCategory) error {
	return db.GORM.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).Updates(filmCategory).Error
}

func DeleteOneFilmCategory(filmCategory models.FilmCategory) error {
	return db.GORM.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).Delete(&filmCategory).Error
}
