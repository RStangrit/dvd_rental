package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/user", PostUserHandler)
	server.GET("/users", GetUsersHandler)
	server.GET("/user/:id", GetUserHandler)
	server.PUT("/user/:id", PutUserHandler)
	server.DELETE("/user/:id", DeleteUserHandler)
}
