package language

import (
	"github.com/gin-gonic/gin"
)

func RegisterLanguageRoutes(server *gin.Engine) {
	server.POST("/language", PostLanguageHandler)
	server.POST("/languages", PostLanguagesHandler)
	server.GET("/languages", GetLanguagesHandler)
	server.GET("/language/:id", GetLanguageHandler)
	server.PUT("/language/:id", PutLanguageHandler)
	server.DELETE("/language/:id", DeleteLanguageHandler)
}
