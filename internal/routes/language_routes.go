package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterLanguageRoutes(server *gin.Engine) {
	server.POST("/language", handlers.PostLanguageHandler)
	server.GET("/languages", handlers.GetLanguagesHandler)
	server.GET("/language/:id", handlers.GetLanguageHandler)
	server.PUT("/language/:id", handlers.PutLanguageHandler)
	server.DELETE("/language/:id", handlers.DeleteLanguageHandler)
}
