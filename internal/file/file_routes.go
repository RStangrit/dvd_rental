package file

import (
	"github.com/gin-gonic/gin"
)

func RegisterFileRoutes(server *gin.Engine) {
	server.GET("/files/*filepath", GetFileHandler)
}
