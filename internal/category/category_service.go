package category

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type CategoryService struct {
	repo *CategoryRepository
}

func NewCategoryService(repo *CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (service *CategoryService) CreateCategory(newCategory *Category) error {
	err := ValidateCategory(newCategory)
	if err != nil {
		return err
	} else {
		return service.repo.InsertCategory(newCategory)
	}
}

func (service *CategoryService) ReadAllCategories(pagination db.Pagination) ([]Category, int64, error) {
	addresses, totalRecords, err := service.repo.SelectAllCategories(pagination)
	if err != nil {
		return nil, 0, err
	}
	return addresses, totalRecords, nil
}

func (service *CategoryService) ReadOneCategory(addressId int64) (*Category, error) {
	address, err := service.repo.SelectOneCategory(addressId)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, fmt.Errorf("address not found")
	}
	return address, nil
}

func (service *CategoryService) UpdateOneCategory(address *Category) error {
	err := ValidateCategory(address)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneCategory(*address)
	}
}

func (service *CategoryService) DeleteOneCategory(address *Category) error {
	return service.repo.DeleteOneCategory(*address)
}

func ValidateCategory(category *Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	if len(category.Name) > 25 {
		return errors.New("category name must be less than or equal to 25 characters")
	}

	return nil
}
