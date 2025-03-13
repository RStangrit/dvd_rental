package inventory

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InventoryRoutes struct {
	handler *InventoryHandler
}

func NewInventoryRoutes(db *gorm.DB) *InventoryRoutes {
	repo := NewInventoryRepository(db)
	service := NewInventoryService(repo)
	handler := NewInventoryHandler(service)

	return &InventoryRoutes{handler: handler}
}

func (route *InventoryRoutes) RegisterInventoryRoutes(server *gin.Engine) {
	server.POST("/inventory", route.handler.PostInventoryHandler)
	server.GET("/inventories", route.handler.GetInventoriesHandler)
	server.GET("/inventory/:id", route.handler.GetInventoryHandler)
	server.PUT("/inventory/:id", route.handler.PutInventoryHandler)
	server.DELETE("/inventory/:id", route.handler.DeleteInventoryHandler)
}
