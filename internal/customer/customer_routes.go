package customer

import (
	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(server *gin.Engine) {
	server.POST("/customer", PostCustomerHandler)
	server.GET("/customers", GetCustomersHandler)
	server.GET("/customer/:id", GetCustomerHandler)
	server.PUT("/customer/:id", PutCustomerHandler)
	server.DELETE("/customer/:id", DeleteCustomerHandler)
}
