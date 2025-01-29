package actor

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postActorHandler(context *gin.Context) {
	var newActor Actor
	var err error

	if err = context.ShouldBindJSON(&newActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = newActor.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = newActor.createActor(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newActor})
}

func getActorsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	actors, totalRecords, err := readAllActors(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actors, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func getActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := readOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actor})
}

func putActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := readOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var updatedActor Actor
	err = context.ShouldBindJSON(&updatedActor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor data format"})
		return
	}

	updatedActor.ActorID = int(actor.ActorID)

	err = updatedActor.updateoneActor()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedActor})
}

func deleteActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := readOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}

	err = actor.deleteOneActor()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": actor})
}
