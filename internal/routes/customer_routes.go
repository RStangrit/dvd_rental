package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(server *gin.Engine) {
	server.POST("/customer", handlers.PostCustomerHandler)
	server.GET("/customers", handlers.GetCustomersHandler)
	server.GET("/customer/:id", handlers.GetCustomerHandler)
	server.PUT("/customer/:id", handlers.PutCustomerHandler)
	server.DELETE("/customer/:id", handlers.DeleteCustomerHandler)
}
