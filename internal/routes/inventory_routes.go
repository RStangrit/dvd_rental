package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(server *gin.Engine) {
	server.POST("/inventory", handlers.PostInventoryHandler)
	server.GET("/inventories", handlers.GetInventoriesHandler)
	server.GET("/inventory/:id", handlers.GetInventoryHandler)
	server.PUT("/inventory/:id", handlers.PutInventoryHandler)
	server.DELETE("/inventory/:id", handlers.DeleteInventoryHandler)
}
