package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFilmActorHandler(context *gin.Context) {
	var newFilmActor models.FilmActor
	var err error

	if err = context.ShouldBindJSON(&newFilmActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateFilmActor(&newFilmActor); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilmActor})
}

func GetFilmsActorsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	filmActors, totalRecords, err := repositories.ReadAllFilmActors(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmActors, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetFilmActorHandler(context *gin.Context) {
	actorID, err := utils.GetIntParam(context, "actor_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor_id ID format"})
		return
	}

	filmID, err := utils.GetIntParam(context, "film_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_id ID format"})
		return
	}

	filmActor, err := repositories.ReadOneFilmActor(actorID, filmID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmActor})
}

func PutFilmActorHandler(context *gin.Context) {
	actorID, err := utils.GetIntParam(context, "actor_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor_id ID format"})
		return
	}

	filmID, err := utils.GetIntParam(context, "film_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_id ID format"})
		return
	}

	_, err = repositories.ReadOneFilmActor(actorID, filmID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFilmActor models.FilmActor
	err = context.ShouldBindJSON(&updatedFilmActor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_actor data format"})
		return
	}

	err = repositories.UpdateOneFilmActor(updatedFilmActor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film_actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilmActor})
}

func DeleteFilmActorHandler(context *gin.Context) {
	actorID, err := utils.GetIntParam(context, "actor_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor_id ID format"})
		return
	}

	filmID, err := utils.GetIntParam(context, "film_id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_id ID format"})
		return
	}

	filmActor, err := repositories.ReadOneFilmActor(actorID, filmID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = repositories.DeleteOneFilmActor(*filmActor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film_actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": filmActor})
}
