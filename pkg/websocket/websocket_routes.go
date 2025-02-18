package websocket

import "github.com/gin-gonic/gin"

func RegisterWSRoutes(server *gin.Engine) {
	server.GET("/ws", websocketHandler)
}
