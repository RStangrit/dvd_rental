package websocket

import (
	"github.com/gin-gonic/gin"
)

type WebSocketRoutes struct {
	cHandler *CentrifugeWebSocketHandler
	mHandler *MelodyWebSocketHandler
	gHandler *GorillaWebSocketHandler
}

func NewWebSocketRoutes() (*WebSocketRoutes, error) {
	cHandler, err := NewCentrifugeWebSocketHandler()
	if err != nil {
		return nil, err
	}

	return &WebSocketRoutes{
		cHandler: cHandler,
		mHandler: NewMelodyWebSocketHandler(),
		gHandler: NewGorillaWebSocketHandler(),
	}, nil
}

func (route *WebSocketRoutes) RegisterWSRoutes(server *gin.Engine) {
	server.GET("/gorilla-ws", route.gHandler.Handle())
	server.GET("/melody-ws", route.mHandler.Handle())
	server.GET("/centrifugo-ws", route.cHandler.Handle())
}
