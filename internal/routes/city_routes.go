package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCityRoutes(server *gin.Engine) {
	server.POST("/city", handlers.PostCityHandler)
	server.GET("/cities", handlers.GetCitiesHandler)
	server.GET("/city/:id", handlers.GetCityHandler)
	server.PUT("/city/:id", handlers.PutCityHandler)
	server.DELETE("/city/:id", handlers.DeleteCityHandler)
}
