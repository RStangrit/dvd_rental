package server

import (
	"main/internal/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	server := gin.Default()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routes.RegisterActorRoutes(server)
	routes.RegisterFilmRoutes(server)
	server.Run()
}
