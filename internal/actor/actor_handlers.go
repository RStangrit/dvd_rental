package actor

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActorHandler struct {
	service *ActorService
}

func NewActorHandler(service *ActorService) *ActorHandler {
	return &ActorHandler{service: service}
}

func (handler *ActorHandler) PostActorHandler(context *gin.Context) {
	var newActor Actor
	var err error

	if err = context.ShouldBindJSON(&newActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateActor(&newActor); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newActor})
}

func (handler *ActorHandler) PostActorsHandler(context *gin.Context) {
	var newActors []*Actor
	var err error

	if err = context.ShouldBindJSON(&newActors); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErrors, createdActors, _ := handler.service.CreateActors(newActors)

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": validationErrors,
			"data":   createdActors,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newActors})
}

func (handler *ActorHandler) GetActorsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	actors, totalRecords, err := handler.service.ReadAllActors(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actors, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *ActorHandler) GetActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := handler.service.ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actor})
}

func (handler *ActorHandler) GetActorFilmsHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actorFilms, err := handler.service.ReadOneActorFilms(actorId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actorFilms})
}

func (handler *ActorHandler) PutActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := handler.service.ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var updatedActor *Actor
	err = context.ShouldBindJSON(&updatedActor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor data format"})
		return
	}

	updatedActor.ActorID = int(actor.ActorID)

	err = handler.service.UpdateOneActor(updatedActor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedActor})
}

func (handler *ActorHandler) DeleteActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := handler.service.ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}

	err = handler.service.DeleteOneActor(actor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": actor})
}
