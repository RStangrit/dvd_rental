package film

import (
	"github.com/gin-gonic/gin"
)

func RegisterFilmRoutes(server *gin.Engine) {
	server.POST("/film", PostFilmHandler)
	server.POST("/films", PostFilmsHandler)
	server.GET("/films", GetFilmshandler)
	server.GET("/film/:id", GetFilmHandler)
	server.GET("/film/:id/actors", GetFilmActorsHandler)
	server.PUT("/film/:id", PutFilmHandler)
	server.DELETE("/film/:id", DeleteFilmHandler)
	server.POST("/film/:id/discount", PostFilmDiscountHandler)
}
