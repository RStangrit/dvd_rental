package websocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type GorillaWebSocketHandler struct {
	upgrader  websocket.Upgrader
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	startOnce sync.Once
	mu        sync.Mutex
}

func NewGorillaWebSocketHandler() *GorillaWebSocketHandler {
	return &GorillaWebSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (handler *GorillaWebSocketHandler) Handle() gin.HandlerFunc {
	return func(context *gin.Context) {
		handler.startOnce.Do(func() {
			go handler.broadcaster()
		})

		conn, err := handler.upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			fmt.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

		handler.mu.Lock()
		handler.clients[conn] = true
		handler.mu.Unlock()

		fmt.Println("Connection has been established")

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Reading error:", err)
				handler.mu.Lock()
				delete(handler.clients, conn)
				handler.mu.Unlock()
				break
			}
			handler.broadcast <- msg
		}
	}
}

func (handler *GorillaWebSocketHandler) broadcaster() {
	for msg := range handler.broadcast {
		handler.mu.Lock()
		for client := range handler.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(handler.clients, client)
			}
		}
		handler.mu.Unlock()
	}
}
