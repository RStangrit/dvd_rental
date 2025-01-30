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

func PostActorHandler(context *gin.Context) {
	var newActor models.Actor
	var err error

	if err = context.ShouldBindJSON(&newActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = services.ValidateActor(newActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateActor(&newActor); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newActor})
}

func GetActorsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	actors, totalRecords, err := repositories.ReadAllActors(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actors, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := repositories.ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actor})
}

func PutActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := repositories.ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var updatedActor models.Actor
	err = context.ShouldBindJSON(&updatedActor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor data format"})
		return
	}

	updatedActor.ActorID = int(actor.ActorID)

	err = repositories.UpdateOneActor(updatedActor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedActor})
}

func DeleteActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := repositories.ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}

	err = repositories.DeleteOneActor(*actor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": actor})
}
