package store

import (
	"github.com/gin-gonic/gin"
)

func RegisterStoreRoutes(server *gin.Engine) {
	server.POST("/store", PostStoreHandler)
	server.GET("/stores", GetStoresHandler)
	server.GET("/store/:id", GetStoreHandler)
	server.PUT("/store/:id", PutStoreHandler)
	server.DELETE("/store/:id", DeleteStoreHandler)
}
