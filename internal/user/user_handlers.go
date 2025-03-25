package user

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) PostUserHandler(context *gin.Context) {
	var newUser User
	var err error

	if err = context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.ValidateUser(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateUser(&newUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newUser})
}

func (handler *UserHandler) GetUsersHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	users, totalRecords, err := handler.service.ReadAllUsers(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": users, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *UserHandler) GetUserHandler(context *gin.Context) {
	userId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := handler.service.ReadOneUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
}

func (handler *UserHandler) PutUserHandler(context *gin.Context) {
	userId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := handler.service.ReadOneUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedUser *User
	err = context.ShouldBindJSON(&updatedUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data format"})
		return
	}

	updatedUser.UserID = user.UserID

	err = handler.service.UpdateOneUser(updatedUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func (handler *UserHandler) DeleteUserHandler(context *gin.Context) {
	userId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := handler.service.ReadOneUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": user})
}

func (handler *UserHandler) LoginUserHandler(context *gin.Context) {
	var inputUser User

	if err := context.ShouldBindJSON(&inputUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := handler.service.LoginUser(inputUser.Email, inputUser.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "successfully authorized", "token": token})
}

func (handler *UserHandler) LogoutUserHandler(context *gin.Context) {
	token := context.GetHeader("Authorization")

	err := handler.service.LogoutUser(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully logged out"})
}
