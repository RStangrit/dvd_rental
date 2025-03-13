package film

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FilmRoutes struct {
	handler *FilmHandler
}

func NewFilmRoutes(db *gorm.DB) *FilmRoutes {
	repo := NewFilmRepository(db)
	service := NewFilmService(repo)
	handler := NewFilmHandler(service)

	return &FilmRoutes{handler: handler}
}

func (route *FilmRoutes) RegisterFilmRoutes(server *gin.Engine) {
	server.POST("/film", route.handler.PostFilmHandler)
	server.POST("/films", route.handler.PostFilmsHandler)
	server.GET("/films", route.handler.GetFilmsHandler)
	server.GET("/film/:id", route.handler.GetFilmHandler)
	server.GET("/film/:id/actors", route.handler.GetFilmActorsHandler)
	server.PUT("/film/:id", route.handler.PutFilmHandler)
	server.DELETE("/film/:id", route.handler.DeleteFilmHandler)
	server.POST("/film/:id/discount", route.handler.PostFilmDiscountHandler)
}
