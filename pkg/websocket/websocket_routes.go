package websocket

import "github.com/gin-gonic/gin"

func RegisterWSRoutes(server *gin.Engine) {
	server.GET("/gorilla-ws", gorillaWebsocketHandler)
	server.GET("/melody-ws", melodyWebsocketHandler)
	server.GET("/centrifugo-ws", centrifugoWebsocketHandler)
}
