package inventory

import (
	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(server *gin.Engine) {
	server.POST("/inventory", PostInventoryHandler)
	server.GET("/inventories", GetInventoriesHandler)
	server.GET("/inventory/:id", GetInventoryHandler)
	server.PUT("/inventory/:id", PutInventoryHandler)
	server.DELETE("/inventory/:id", DeleteInventoryHandler)
}
