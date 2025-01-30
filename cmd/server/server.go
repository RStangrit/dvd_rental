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
	routes.RegisterLanguageRoutes(server)
	routes.RegisterActorRoutes(server)
	routes.RegisterFilmRoutes(server)
	routes.RegisterCategoryRoutes(server)
	routes.RegisterFilmActorRoutes(server)
	routes.RegisterInventoryRoutes(server)
	routes.RegisterFilmCategoryRoutes(server)
	routes.RegisterCountryRoutes(server)
	server.Run()
}
