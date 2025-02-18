package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader is needed to upgrade the HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, //Allow any connections
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
	startOnce sync.Once // broadcaster() will be launched only once
)

func websocketHandler(context *gin.Context) {

	startOnce.Do(func() {
		go broadcaster()
	})

	// Updating the connection to WebSocket
	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	log.Println("Connection has been established")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Reading error:", err)
			delete(clients, conn)
			break
		}
		broadcast <- msg
	}
}

func broadcaster() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
