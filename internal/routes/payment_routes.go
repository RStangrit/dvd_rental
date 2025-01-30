package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(server *gin.Engine) {
	server.POST("/payment", handlers.PostPaymentHandler)
	server.GET("/payments", handlers.GetPaymentsHandler)
	server.GET("/payment/:id", handlers.GetPaymentHandler)
	server.PUT("/payment/:id", handlers.PutPaymentHandler)
	server.DELETE("/payment/:id", handlers.DeletePaymentHandler)
}
