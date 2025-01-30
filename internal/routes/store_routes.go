package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterStoreRoutes(server *gin.Engine) {
	server.POST("/store", handlers.PostStoreHandler)
	server.GET("/stores", handlers.GetStoresHandler)
	server.GET("/store/:id", handlers.GetStoreHandler)
	server.PUT("/store/:id", handlers.PutStoreHandler)
	server.DELETE("/store/:id", handlers.DeleteStoreHandler)
}
