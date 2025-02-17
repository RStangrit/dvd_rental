package user

import (
	"main/pkg/auth"
	"main/pkg/db"
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

	if err = CreateUser(db.GORM, &newUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newUser})
}

func GetUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}

func GetUsersHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	users, totalRecords, err := ReadAllUsers(db.GORM, pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": users, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func PutUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}

func DeleteUserHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "handler in development"})
}

func LoginUserHandler(context *gin.Context) {
	var inputUser User
	var err error

	if err = context.ShouldBindJSON(&inputUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ReadOneUserByEmail(db.GORM, inputUser.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = auth.CompareHashAndPassword(user.Password, inputUser.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.CreateToken(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "successfully authorized", "token": tokenString})
}
