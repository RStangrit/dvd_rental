package city

import (
	"github.com/gin-gonic/gin"
)

func RegisterCityRoutes(server *gin.Engine) {
	server.POST("/city", PostCityHandler)
	server.GET("/cities", GetCitiesHandler)
	server.GET("/city/:id", GetCityHandler)
	server.PUT("/city/:id", PutCityHandler)
	server.DELETE("/city/:id", DeleteCityHandler)
}
