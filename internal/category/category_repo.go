package category

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) InsertCategory(newCategory *Category) error {
	return repo.db.Table("category").Create(&newCategory).Error
}

func (repo *CategoryRepository) SelectAllCategories(pagination db.Pagination) ([]Category, int64, error) {
	var categories []Category
	var totalRecords int64

	repo.db.Table("category").Count(&totalRecords)
	err := repo.db.Table("category").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("category_id asc").Find(&categories).Error
	return categories, totalRecords, err
}

func (repo *CategoryRepository) SelectOneCategory(categoryId int64) (*Category, error) {
	var category Category
	err := repo.db.Table("category").First(&category, categoryId).Error
	return &category, err
}

func (repo *CategoryRepository) UpdateOneCategory(category Category) error {
	return repo.db.Table("category").Omit("category_id").Updates(category).Error
}

func (repo *CategoryRepository) DeleteOneCategory(category Category) error {
	return repo.db.Delete(&category).Error
}
