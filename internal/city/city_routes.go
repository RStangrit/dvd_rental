package city

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CityRoutes struct {
	handler *CityHandler
}

func NewCityRoutes(db *gorm.DB) *CityRoutes {
	repo := NewCityRepository(db)
	service := NewCityService(repo)
	handler := NewCityHandler(service)

	return &CityRoutes{handler: handler}
}

func (route *CityRoutes) RegisterCityRoutes(server *gin.Engine) {
	server.POST("/city", route.handler.PostCityHandler)
	server.GET("/cities", route.handler.GetCitiesHandler)
	server.GET("/city/:id", route.handler.GetCityHandler)
	server.PUT("/city/:id", route.handler.PutCityHandler)
	server.DELETE("/city/:id", route.handler.DeleteCityHandler)
}
