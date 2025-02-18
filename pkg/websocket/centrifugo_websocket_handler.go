package websocket

import (
	"log"

	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
)

var node *centrifuge.Node

func init() {
	var err error
	node, err = centrifuge.New(centrifuge.Config{})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := node.Run(); err != nil {
			log.Fatal(err)
		}
	}()
}

func centrifugoWebsocketHandler(context *gin.Context) {
	if node == nil {
		context.JSON(500, gin.H{"error": "Centrifuge server is not running"})
		return
	}

	wsHandler := centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{})
	wsHandler.ServeHTTP(context.Writer, context.Request)
}
