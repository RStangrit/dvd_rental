package film_category

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateFilmCategory(db *gorm.DB, newFilmCategory *FilmCategory) error {
	return db.Table("film_category").Create(&newFilmCategory).Error
}

func ReadAllFilmCategories(db *gorm.DB, pagination db.Pagination, sortParams string) ([]FilmCategory, int64, error) {
	var filmCategories []FilmCategory
	var totalRecords int64

	db.Table("film_category").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("film_category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(sortParams).Find(&filmCategories).Error
	return filmCategories, totalRecords, err
}

func ReadOneFilmCategory(db *gorm.DB, filmID, categoryID int64) (*FilmCategory, error) {
	var filmCategory FilmCategory
	err := db.Table("film_category").Where("film_id = ? AND category_id = ?", filmID, categoryID).First(&filmCategory).Error
	return &filmCategory, err
}

func UpdateOneFilmCategory(db *gorm.DB, filmCategory FilmCategory) error {
	return db.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).Updates(filmCategory).Error
}

func DeleteOneFilmCategory(db *gorm.DB, filmCategory FilmCategory) error {
	return db.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).Delete(&filmCategory).Error
}
