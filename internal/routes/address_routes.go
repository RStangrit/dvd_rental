package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAddressRoutes(server *gin.Engine) {
	server.POST("/address", handlers.PostAddressHandler)
	server.GET("/addresses", handlers.GetAddressesHandler)
	server.GET("/address/:id", handlers.GetAddressHandler)
	server.PUT("/address/:id", handlers.PutAddressHandler)
	server.DELETE("/address/:id", handlers.DeleteAddressHandler)
}
