package payment

import (
	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(server *gin.Engine) {
	server.POST("/payment", PostPaymentHandler)
	server.GET("/payments", GetPaymentsHandler)
	server.GET("/payment/:id", GetPaymentHandler)
	server.PUT("/payment/:id", PutPaymentHandler)
	server.DELETE("/payment/:id", DeletePaymentHandler)
}
