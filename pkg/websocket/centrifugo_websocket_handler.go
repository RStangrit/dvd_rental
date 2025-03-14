package websocket

import (
	"log"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
)

type CentrifugeWebSocketHandler struct {
	node *centrifuge.Node
}

func NewCentrifugeWebSocketHandler() (*CentrifugeWebSocketHandler, error) {
	node, err := centrifuge.New(centrifuge.Config{})
	if err != nil {
		return nil, err
	}

	go func() {
		if err := node.Run(); err != nil {
			log.Fatal("Centrifuge error:", err)
		}
	}()

	return &CentrifugeWebSocketHandler{node: node}, nil
}

func (h *CentrifugeWebSocketHandler) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		if h.node == nil {
			log.Println("Centrifuge server is not running")
			context.JSON(500, gin.H{"error": "Centrifuge server is not running"})
			return
		}

		wsHandler := centrifuge.NewWebsocketHandler(h.node, centrifuge.WebsocketConfig{})
		wsHandler.ServeHTTP(context.Writer, context.Request)
	}
}

func (h *CentrifugeWebSocketHandler) Shutdown(context *gin.Context) {
	if h.node != nil {
		h.node.Shutdown(context)
		log.Println("Centrifuge server stopped")
	}
}
