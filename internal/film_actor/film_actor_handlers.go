package film_actor

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilmActorHandler struct {
	service *FilmActorService
}

func NewFilmActorHandler(service *FilmActorService) *FilmActorHandler {
	return &FilmActorHandler{service: service}
}

func (handler *FilmActorHandler) PostFilmActorHandler(context *gin.Context) {
	var newFilmActor FilmActor
	var err error

	if err = context.ShouldBindJSON(&newFilmActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateFilmActor(&newFilmActor); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newFilmActor})
}

func (handler *FilmActorHandler) GetFilmsActorsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	filmActors, totalRecords, err := handler.service.ReadAllFilmActors(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmActors, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *FilmActorHandler) GetFilmActorHandler(context *gin.Context) {
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

	filmActor, err := handler.service.ReadOneFilmActor(actorID, filmID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": filmActor})
}

func (handler *FilmActorHandler) PutFilmActorHandler(context *gin.Context) {
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

	_, err = handler.service.ReadOneFilmActor(actorID, filmID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedFilmActor FilmActor
	err = context.ShouldBindJSON(&updatedFilmActor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid film_actor data format"})
		return
	}

	err = handler.service.UpdateOneFilmActor(int(actorID), int(filmID), &updatedFilmActor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film_actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedFilmActor})
}

func (handler *FilmActorHandler) DeleteFilmActorHandler(context *gin.Context) {
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

	filmActor, err := handler.service.ReadOneFilmActor(actorID, filmID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneFilmActor(filmActor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film_actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": filmActor})
}
