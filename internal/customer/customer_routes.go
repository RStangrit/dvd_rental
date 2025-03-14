package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRoutes struct {
	handler *CustomerHandler
}

func NewCustomerRoutes(db *gorm.DB) *CustomerRoutes {
	repo := NewCustomerRepository(db)
	service := NewCustomerService(repo)
	handler := NewCustomerHandler(service)

	return &CustomerRoutes{handler: handler}
}

func (route *CustomerRoutes) RegisterCustomerRoutes(server *gin.Engine) {
	server.POST("/customer", route.handler.PostCustomerHandler)
	server.GET("/customers", route.handler.GetCustomersHandler)
	server.GET("/customer/:id", route.handler.GetCustomerHandler)
	server.PUT("/customer/:id", route.handler.PutCustomerHandler)
	server.DELETE("/customer/:id", route.handler.DeleteCustomerHandler)
}
