package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterFilmRoutes(server *gin.Engine) {
	server.POST("/film", handlers.PostFilmHandler)
	server.GET("/films", handlers.GetFilmshandler)
	server.GET("/film/:id", handlers.GetFilmHandler)
	server.PUT("/film/:id", handlers.PutFilmHandler)
	server.DELETE("/film/:id", handlers.DeleteFilmHandler)
}
