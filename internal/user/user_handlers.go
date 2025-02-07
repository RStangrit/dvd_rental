package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUserHandler(context *gin.Context) {
	var newUser User
	var err error

	if err = context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateUser(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateUser(&newUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newUser})
}
func GetUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}
func GetUsersHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}
func PutUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}
func DeleteUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}
