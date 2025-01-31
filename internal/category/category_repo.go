package category

import (
	"main/pkg/db"
)

func CreateCategory(newCategory *Category) error {
	return db.GORM.Table("category").Create(&newCategory).Error
}

func ReadAllCategories(pagination db.Pagination) ([]Category, int64, error) {
	var categories []Category
	var totalRecords int64

	db.GORM.Table("category").Count(&totalRecords)
	err := db.GORM.Table("category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("category_id asc").Find(&categories).Error
	return categories, totalRecords, err
}

func ReadOneCategory(categoryId int64) (*Category, error) {
	var category Category
	err := db.GORM.Table("category").First(&category, categoryId).Error
	return &category, err
}

func UpdateOneCategory(category Category) error {
	return db.GORM.Table("category").Omit("category_id").Updates(category).Error
}

func DeleteOneCategory(category Category) error {
	return db.GORM.Delete(&category).Error
}
