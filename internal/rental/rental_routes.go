package rental

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RentalRoutes struct {
	handler *RentalHandler
}

func NewRentalRoutes(db *gorm.DB) *RentalRoutes {
	repo := NewRentalRepository(db)
	service := NewRentalService(repo)
	handler := NewRentalHandler(service)

	return &RentalRoutes{handler: handler}
}

func (route *RentalRoutes) RegisterRentalRoutes(server *gin.Engine) {
	server.POST("/rental", route.handler.PostRentalHandler)
	server.GET("/rentals", route.handler.GetRentalsHandler)
	server.GET("/rental/:id", route.handler.GetRentalHandler)
	server.PUT("/rental/:id", route.handler.PutRentalHandler)
	server.DELETE("/rental/:id", route.handler.DeleteRentalHandler)
}
