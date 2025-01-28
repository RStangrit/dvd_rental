package server

import (
	"main/internal/actor"
	"main/internal/film"
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
	actor.RegisterRoutes(server)
	film.RegisterRoutes(server)
	server.Run()
}
