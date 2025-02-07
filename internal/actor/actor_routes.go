package actor

import (
	"github.com/gin-gonic/gin"
)

func RegisterActorRoutes(server *gin.Engine) {
	server.POST("/actor", PostActorHandler)
	server.POST("/actors", PostActorsHandler)
	server.GET("/actors", GetActorsHandler)
	server.GET("/actor/:id", GetActorHandler)
	server.GET("/actor/:id/films", GetActorFilmsHandler)
	server.PUT("/actor/:id", PutActorHandler)
	server.DELETE("/actor/:id", DeleteActorHandler)
}
