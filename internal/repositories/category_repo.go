package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateCategory(newCategory *models.Category) error {
	return db.GORM.Table("category").Create(&newCategory).Error
}

func ReadAllCategories(pagination db.Pagination) ([]models.Category, int64, error) {
	var categories []models.Category
	var totalRecords int64

	db.GORM.Table("category").Count(&totalRecords)
	err := db.GORM.Table("category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("category_id asc").Find(&categories).Error
	return categories, totalRecords, err
}

func ReadOneCategory(categoryId int64) (*models.Category, error) {
	var category models.Category
	err := db.GORM.Table("category").First(&category, categoryId).Error
	return &category, err
}

func UpdateOneCategory(category models.Category) error {
	return db.GORM.Table("category").Omit("category_id").Updates(category).Error
}

func DeleteOneCategory(category models.Category) error {
	return db.GORM.Delete(&category).Error
}
