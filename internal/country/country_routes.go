package country

import (
	"github.com/gin-gonic/gin"
)

func RegisterCountryRoutes(server *gin.Engine) {
	server.POST("/country", PostCountryHandler)
	server.GET("/countries", GetCountriesHandler)
	server.GET("/country/:id", GetCountryHandler)
	server.PUT("/country/:id", PutCountryHandler)
	server.DELETE("/country/:id", DeleteCountryHandler)
}
