package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var (
	m = melody.New()
)

func init() {
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastOthers(msg, s)
	})
}

func melodyWebsocketHandler(context *gin.Context) {
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	defer m.Close()

	m.HandleRequest(context.Writer, context.Request)
}
