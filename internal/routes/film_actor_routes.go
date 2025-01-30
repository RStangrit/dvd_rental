package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterFilmActorRoutes(server *gin.Engine) {
	server.POST("/film-actor", handlers.PostFilmActorHandler)
	server.GET("/film-actors", handlers.GetFilmsActorsHandler)
	server.GET("/film_actor/:actor_id/:film_id", handlers.GetFilmActorHandler)
	server.PUT("/film_actor/:actor_id/:film_id", handlers.PutFilmActorHandler)
	server.DELETE("/film_actor/:actor_id/:film_id", handlers.DeleteFilmActorHandler)
}
