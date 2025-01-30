package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFilmCategoryHandler(context *gin.Context) {
	var newFilmCategory models.FilmCategory
	var err error

	if err = context.ShouldBindJSON(&newFilmCategory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateFilmCategory(&newFilmCategory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilmCategory})
}

func GetFilmCategoriesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	filmCategories, totalRecords, err := repositories.ReadAllFilmCategories(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmCategories, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetFilmCategoryHandler(context *gin.Context) {
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

	filmCategory, err := repositories.ReadOneFilmCategory(filmID, categoryID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmCategory})
}

func PutFilmCategoryHandler(context *gin.Context) {
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

	_, err = repositories.ReadOneFilmCategory(filmID, categoryID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFilmCategory models.FilmCategory
	err = context.ShouldBindJSON(&updatedFilmCategory)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_category data format"})
		return
	}

	err = repositories.UpdateOneFilmCategory(updatedFilmCategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film_category"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilmCategory})
}

func DeleteFilmCategoryHandler(context *gin.Context) {
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

	filmCategory, err := repositories.ReadOneFilmCategory(filmID, categoryID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = repositories.DeleteOneFilmCategory(*filmCategory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film_category"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": filmCategory})
}
