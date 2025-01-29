package film

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/film", postFilmHandler)
	server.GET("/films", getFilmshandler)
	server.GET("/film/:id", getFilmHandler)
	server.PUT("/film/:id")
	server.DELETE("/film/:id")
}
