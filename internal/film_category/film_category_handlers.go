package film_category

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilmCategoryHandler struct {
	service *FilmCategoryService
}

func NewFilmCategoryHandler(service *FilmCategoryService) *FilmCategoryHandler {
	return &FilmCategoryHandler{service: service}
}

func (handler *FilmCategoryHandler) PostFilmCategoryHandler(context *gin.Context) {
	var newFilmCategory FilmCategory

	if err := context.ShouldBindJSON(&newFilmCategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.service.CreateFilmCategory(&newFilmCategory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilmCategory})
}

func (handler *FilmCategoryHandler) GetFilmCategoriesHandler(context *gin.Context) {
	var pagination db.Pagination

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	field := context.DefaultQuery("field", "film_id")
	order := context.DefaultQuery("order", "asc")
	sortParams := utils.JoinStrings(field, order)

	filmCategories, totalRecords, err := handler.service.ReadAllFilmCategories(pagination, sortParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmCategories, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *FilmCategoryHandler) GetFilmCategoryHandler(context *gin.Context) {
	filmID, err := utils.GetIntParam(context, "film_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_id format"})
		return
	}

	categoryID, err := utils.GetIntParam(context, "category_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id format"})
		return
	}

	filmCategory, err := handler.service.ReadOneFilmCategory(filmID, categoryID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmCategory})
}

func (handler *FilmCategoryHandler) PutFilmCategoryHandler(context *gin.Context) {
	filmID, err := utils.GetIntParam(context, "film_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_id format"})
		return
	}

	categoryID, err := utils.GetIntParam(context, "category_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id format"})
		return
	}

	_, err = handler.service.ReadOneFilmCategory(filmID, categoryID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFilmCategory FilmCategory
	if err = context.ShouldBindJSON(&updatedFilmCategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_category data format"})
		return
	}

	err = handler.service.UpdateOneFilmCategory(int(filmID), int(categoryID), &updatedFilmCategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film_category"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilmCategory})
}

func (handler *FilmCategoryHandler) DeleteFilmCategoryHandler(context *gin.Context) {
	filmID, err := utils.GetIntParam(context, "film_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_id format"})
		return
	}

	categoryID, err := utils.GetIntParam(context, "category_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id format"})
		return
	}

	filmCategory, err := handler.service.ReadOneFilmCategory(filmID, categoryID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneFilmCategory(filmCategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film_category"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": filmCategory})
}
