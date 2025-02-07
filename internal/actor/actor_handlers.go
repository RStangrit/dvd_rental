package actor

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostActorHandler(context *gin.Context) {
	var newActor Actor
	var err error

	if err = context.ShouldBindJSON(&newActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateActor(&newActor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateActor(&newActor); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newActor})
}

func PostActorsHandler(context *gin.Context) {
	var newACtors []Actor
	var err error

	if err = context.ShouldBindJSON(&newACtors); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var validationErrors []string
	var createdActors []Actor

	for _, newActor := range newACtors {
		if err = ValidateActor(&newActor); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err = CreateActor(&newActor); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		createdActors = append(createdActors, newActor)
	}

	if len(validationErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": validationErrors,
			"data":   createdActors,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newACtors})

}

func GetActorsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	actors, totalRecords, err := ReadAllActors(pagination)
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

	actor, err := ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actor})
}

func GetActorFilmsHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actorFilms, err := ReadOneActorFilms(actorId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": actorFilms})
}

func PutActorHandler(context *gin.Context) {
	actorId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid actor ID format"})
		return
	}

	actor, err := ReadOneActor(actorId)
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

	err = UpdateOneActor(updatedActor)
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

	actor, err := ReadOneActor(actorId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Actor not found"})
		return
	}

	err = DeleteOneActor(*actor)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete actor"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": actor})
}
