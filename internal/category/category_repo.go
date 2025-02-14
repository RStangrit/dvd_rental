package category

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, newCategory *Category) error {
	return db.Table("category").Create(&newCategory).Error
}

func ReadAllCategories(db *gorm.DB, pagination db.Pagination) ([]Category, int64, error) {
	var categories []Category
	var totalRecords int64

	db.Table("category").Count(&totalRecords)
	err := db.Table("category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("category_id asc").Find(&categories).Error
	return categories, totalRecords, err
}

func ReadOneCategory(db *gorm.DB, categoryId int64) (*Category, error) {
	var category Category
	err := db.Table("category").First(&category, categoryId).Error
	return &category, err
}

func UpdateOneCategory(db *gorm.DB, category Category) error {
	return db.Table("category").Omit("category_id").Updates(category).Error
}

func DeleteOneCategory(db *gorm.DB, category Category) error {
	return db.Delete(&category).Error
}
