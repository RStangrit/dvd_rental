package address

import (
	"github.com/gin-gonic/gin"
)

func RegisterAddressRoutes(server *gin.Engine) {
	server.POST("/address", PostAddressHandler)
	server.GET("/addresses", GetAddressesHandler)
	server.GET("/address/:id", GetAddressHandler)
	server.PUT("/address/:id", PutAddressHandler)
	server.DELETE("/address/:id", DeleteAddressHandler)
}
