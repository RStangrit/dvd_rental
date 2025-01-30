package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/internal/services"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFilmHandler(context *gin.Context) {
	var newFilm models.Film
	var err error

	if err = context.ShouldBindJSON(&newFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = services.ValidateFilm(&newFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateFilm(&newFilm); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilm})
}

func GetFilmshandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	films, totalRecords, err := repositories.ReadAllFilms(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": films, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := repositories.ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": film})
}

func PutFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := repositories.ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFilm models.Film
	err = context.ShouldBindJSON(&updatedFilm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film data format"})
		return
	}

	if err = services.ValidateFilm(&updatedFilm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFilm.FilmID = int(film.FilmID)

	err = repositories.UpdateOneFilm(updatedFilm)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilm})
}

func DeleteFilmHandler(context *gin.Context) {
	filmId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film ID format"})
		return
	}

	film, err := repositories.ReadOneFilm(filmId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = repositories.DeleteOneFilm(*film)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": film})
}
