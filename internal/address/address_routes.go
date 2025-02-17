package address

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddressRoutes struct {
	handler *AddressHandler
}

func NewAddressRoutes(db *gorm.DB) *AddressRoutes {
	repo := NewAddressRepository(db)
	service := NewAddressService(repo)
	handler := NewAddressHandler(service)

	return &AddressRoutes{handler: handler}
}

func (route *AddressRoutes) RegisterAddressRoutes(server *gin.Engine) {
	server.POST("/address", route.handler.PostAddressHandler)
	server.GET("/addresses", route.handler.GetAddressesHandler)
	server.GET("/address/:id", route.handler.GetAddressHandler)
	server.PUT("/address/:id", route.handler.PutAddressHandler)
	server.DELETE("/address/:id", route.handler.DeleteAddressHandler)
}
