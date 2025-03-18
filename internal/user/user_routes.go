package user

import (
	"main/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRoutes struct {
	handler *UserHandler
}

func NewUserRoutes(db *gorm.DB) *UserRoutes {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	return &UserRoutes{handler: handler}
}

func (route *UserRoutes) RegisterUserRoutes(server *gin.Engine) {
	server.POST("/user", route.handler.PostUserHandler)
	server.GET("/users", middleware.AuthMiddleware(), route.handler.GetUsersHandler)
	server.GET("/user/:id", middleware.AuthMiddleware(), route.handler.GetUserHandler)
	server.PUT("/user/:id", middleware.AuthMiddleware(), route.handler.PutUserHandler)
	server.DELETE("/user/:id", middleware.AuthMiddleware(), route.handler.DeleteUserHandler)
	server.POST("/login", route.handler.LoginUserHandler)
}
