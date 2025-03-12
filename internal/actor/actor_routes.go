package actor

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActorRoutes struct {
	handler *ActorHandler
}

func NewActorRoutes(db *gorm.DB) *ActorRoutes {
	repo := NewActorRepository(db)
	service := NewActorService(repo)
	handler := NewActorHandler(service)

	return &ActorRoutes{handler: handler}
}

func (route *ActorRoutes) RegisterActorRoutes(server *gin.Engine) {
	server.POST("/actor", route.handler.PostActorHandler)
	server.POST("/actors", route.handler.PostActorsHandler)
	server.GET("/actors", route.handler.GetActorsHandler)
	server.GET("/actor/:id", route.handler.GetActorHandler)
	server.GET("/actor/:id/films", route.handler.GetActorFilmsHandler)
	server.PUT("/actor/:id", route.handler.PutActorHandler)
	server.DELETE("/actor/:id", route.handler.DeleteActorHandler)
}
