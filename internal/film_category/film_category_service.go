package film_category

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type FilmCategoryService struct {
	repo *FilmCategoryRepository
}

func NewFilmCategoryService(repo *FilmCategoryRepository) *FilmCategoryService {
	return &FilmCategoryService{repo: repo}
}

func (service *FilmCategoryService) CreateFilmCategory(newFilmCategory *FilmCategory) error {
	if err := service.ValidateFilmCategory(newFilmCategory); err != nil {
		return err
	}
	return service.repo.InsertFilmCategory(newFilmCategory)
}

func (service *FilmCategoryService) ReadAllFilmCategories(pagination db.Pagination, sortParams string) ([]FilmCategory, int64, error) {
	filmCategories, totalRecords, err := service.repo.SelectAllFilmCategories(pagination, sortParams)
	if err != nil {
		return nil, 0, err
	}
	return filmCategories, totalRecords, nil
}

func (service *FilmCategoryService) ReadOneFilmCategory(filmID, categoryID int64) (*FilmCategory, error) {
	filmCategory, err := service.repo.SelectOneFilmCategory(filmID, categoryID)
	if err != nil {
		return nil, err
	}
	if filmCategory == nil {
		return nil, fmt.Errorf("filmCategory not found")
	}
	return filmCategory, nil
}

func (service *FilmCategoryService) UpdateOneFilmCategory(filmID, categoryID int, updatedFilmCategory *FilmCategory) error {
	if err := service.ValidateFilmCategory(updatedFilmCategory); err != nil {
		return err
	}
	return service.repo.UpdateOneFilmCategory(filmID, categoryID, updatedFilmCategory)
}

func (service *FilmCategoryService) DeleteOneFilmCategory(filmCategory *FilmCategory) error {
	return service.repo.DeleteOneFilmCategory(*filmCategory)
}

func (service *FilmCategoryService) ValidateFilmCategory(filmCategory *FilmCategory) error {
	if filmCategory.FilmID <= 0 {
		return errors.New("film_id must be a positive integer")
	}

	if filmCategory.CategoryID <= 0 {
		return errors.New("category_id must be a positive integer")
	}

	return nil
}
