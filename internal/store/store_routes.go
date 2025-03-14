package store

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreRoutes struct {
	handler *StoreHandler
}

func NewStoreRoutes(db *gorm.DB) *StoreRoutes {
	repo := NewStoreRepository(db)
	service := NewStoreService(repo)
	handler := NewStoreHandler(service)

	return &StoreRoutes{handler: handler}
}

func (route *StoreRoutes) RegisterStoreRoutes(server *gin.Engine) {
	server.POST("/store", route.handler.PostStoreHandler)
	server.GET("/stores", route.handler.GetStoresHandler)
	server.GET("/store/:id", route.handler.GetStoreHandler)
	server.PUT("/store/:id", route.handler.PutStoreHandler)
	server.DELETE("/store/:id", route.handler.DeleteStoreHandler)
}
