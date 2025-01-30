package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRentalRoutes(server *gin.Engine) {
	server.POST("/rental", handlers.PostRentalHandler)
	server.GET("/rentals", handlers.GetRentalsHandler)
	server.GET("/rental/:id", handlers.GetRentalHandler)
	server.PUT("/rental/:id", handlers.PutRentalHandler)
	server.DELETE("/rental/:id", handlers.DeleteRentalHandler)
}
