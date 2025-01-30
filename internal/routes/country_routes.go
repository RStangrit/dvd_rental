package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCountryRoutes(server *gin.Engine) {
	server.POST("/country", handlers.PostCountryHandler)
	server.GET("/countries", handlers.GetCountriesHandler)
	server.GET("/country/:id", handlers.GetCountryHandler)
	server.PUT("/country/:id", handlers.PutCountryHandler)
	server.DELETE("/country/:id", handlers.DeleteCountryHandler)
}
