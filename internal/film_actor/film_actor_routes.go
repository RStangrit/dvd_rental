package film_actor

import (
	"github.com/gin-gonic/gin"
)

func RegisterFilmActorRoutes(server *gin.Engine) {
	server.POST("/film-actor", PostFilmActorHandler)
	server.GET("/film-actors", GetFilmsActorsHandler)
	server.GET("/film_actor/:actor_id/:film_id", GetFilmActorHandler)
	server.PUT("/film_actor/:actor_id/:film_id", PutFilmActorHandler)
	server.DELETE("/film_actor/:actor_id/:film_id", DeleteFilmActorHandler)
}
