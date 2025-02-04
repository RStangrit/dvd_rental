package development

import (
	"github.com/gin-gonic/gin"
)

func RegisterDevelopmentRoutes(server *gin.Engine) {
	server.GET("/test", GetTestHandler)
}
