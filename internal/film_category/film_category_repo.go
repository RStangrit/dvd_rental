package film_category

import (
	"main/pkg/db"
)

func CreateFilmCategory(newFilmCategory *FilmCategory) error {
	return db.GORM.Table("film_category").Create(&newFilmCategory).Error
}

func ReadAllFilmCategories(pagination db.Pagination, sortParams string) ([]FilmCategory, int64, error) {
	var filmCategories []FilmCategory
	var totalRecords int64

	db.GORM.Table("film_category").Count(&totalRecords)
	err := db.GORM.Table("film_category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(sortParams).Find(&filmCategories).Error
	return filmCategories, totalRecords, err
}

func ReadOneFilmCategory(filmID, categoryID int64) (*FilmCategory, error) {
	var filmCategory FilmCategory
	err := db.GORM.Table("film_category").Where("film_id = ? AND category_id = ?", filmID, categoryID).First(&filmCategory).Error
	return &filmCategory, err
}

func UpdateOneFilmCategory(filmCategory FilmCategory) error {
	return db.GORM.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).Updates(filmCategory).Error
}

func DeleteOneFilmCategory(filmCategory FilmCategory) error {
	return db.GORM.Table("film_category").Where("film_id = ? AND category_id = ?", filmCategory.FilmID, filmCategory.CategoryID).Delete(&filmCategory).Error
}
