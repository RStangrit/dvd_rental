package development

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DevelopmentRoutes struct {
	handler *DevelopmentHandler
}

func NewDevelopmentRoutes(db *gorm.DB) *DevelopmentRoutes {
	repo := NewDevelopmentRepository(db)
	service := NewDevelopmentService(repo)
	handler := NewDevelopmentHandler(service)

	return &DevelopmentRoutes{handler: handler}
}

func (route *DevelopmentRoutes) RegisterDevelopmentRoutes(server *gin.Engine) {
	server.GET("/test", route.handler.GetTestHandler)
	server.POST("/development/transaction", route.handler.GetTestHandler)
}
