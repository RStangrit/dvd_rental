package user

import (
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/user", PostUserHandler)
	server.GET("/users", middleware.AuthMiddleware(), GetUsersHandler)
	server.GET("/user/:id", GetUserHandler)
	server.PUT("/user/:id", PutUserHandler)
	server.DELETE("/user/:id", DeleteUserHandler)
	server.POST("/login", LoginUserHandler)
}
