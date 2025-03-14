package payment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaymentRoutes struct {
	handler *PaymentHandler
}

func NewPaymentRoutes(db *gorm.DB) *PaymentRoutes {
	repo := NewPaymentRepository(db)
	service := NewPaymentService(repo)
	handler := NewPaymentHandler(service)

	return &PaymentRoutes{handler: handler}
}

func (route *PaymentRoutes) RegisterPaymentRoutes(server *gin.Engine) {
	server.POST("/payment", route.handler.PostPaymentHandler)
	server.GET("/payments", route.handler.GetPaymentsHandler)
	server.GET("/payment/:id", route.handler.GetPaymentHandler)
	server.PUT("/payment/:id", route.handler.PutPaymentHandler)
	server.DELETE("/payment/:id", route.handler.DeletePaymentHandler)
}
