package rental

import (
	"github.com/gin-gonic/gin"
)

func RegisterRentalRoutes(server *gin.Engine) {
	server.POST("/rental", PostRentalHandler)
	server.GET("/rentals", GetRentalsHandler)
	server.GET("/rental/:id", GetRentalHandler)
	server.PUT("/rental/:id", PutRentalHandler)
	server.DELETE("/rental/:id", DeleteRentalHandler)
}
