package actor

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/actor", postActorHandler)
	server.GET("/actor", getActorsHandler)
	server.GET("/actor/:id", getActorHandler)
	server.PUT("/actor/:id", putActorHandler)
	server.DELETE("/actor/:id", deleteActorHandler)
}
