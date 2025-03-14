package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type MelodyWebSocketHandler struct {
	m *melody.Melody
}

func NewMelodyWebSocketHandler() *MelodyWebSocketHandler {
	m := melody.New()

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastOthers(msg, s)
	})

	return &MelodyWebSocketHandler{m: m}
}

func (h *MelodyWebSocketHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		h.m.HandleRequest(c.Writer, c.Request)
	}
}
