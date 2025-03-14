package country

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CountryRoutes struct {
	handler *CountryHandler
}

func NewCountryRoutes(db *gorm.DB) *CountryRoutes {
	repo := NewCountryRepository(db)
	service := NewCountryService(repo)
	handler := NewCountryHandler(service)

	return &CountryRoutes{handler: handler}
}

func (route *CountryRoutes) RegisterCountryRoutes(server *gin.Engine) {
	server.POST("/country", route.handler.PostCountryHandler)
	server.GET("/countries", route.handler.GetCountriesHandler)
	server.GET("/country/:id", route.handler.GetCountryHandler)
	server.PUT("/country/:id", route.handler.PutCountryHandler)
	server.DELETE("/country/:id", route.handler.DeleteCountryHandler)
}
