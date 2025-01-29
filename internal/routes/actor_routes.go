package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterActorRoutes(server *gin.Engine) {
	server.POST("/actor", handlers.PostActorHandler)
	server.GET("/actors", handlers.GetActorsHandler)
	server.GET("/actor/:id", handlers.GetActorHandler)
	server.PUT("/actor/:id", handlers.PutActorHandler)
	server.DELETE("/actor/:id", handlers.DeleteActorHandler)
}
