package film_actor

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FilmActorRoutes struct {
	handler *FilmActorHandler
}

func NewFilmActorRoutes(db *gorm.DB) *FilmActorRoutes {
	repo := NewFilmActorRepository(db)
	service := NewFilmActorService(repo)
	handler := NewFilmActorHandler(service)

	return &FilmActorRoutes{handler: handler}
}

func (route *FilmActorRoutes) RegisterFilmActorRoutes(server *gin.Engine) {
	server.POST("/film-actor", route.handler.PostFilmActorHandler)
	server.GET("/film-actors", route.handler.GetFilmsActorsHandler)
	server.GET("/film_actor/:actor_id/:film_id", route.handler.GetFilmActorHandler)
	server.PUT("/film_actor/:actor_id/:film_id", route.handler.PutFilmActorHandler)
	server.DELETE("/film_actor/:actor_id/:film_id", route.handler.DeleteFilmActorHandler)
}
