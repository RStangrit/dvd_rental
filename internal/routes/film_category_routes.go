package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterFilmCategoryRoutes(server *gin.Engine) {
	server.POST("/film-category", handlers.PostFilmCategoryHandler)
	server.GET("/film-categories", handlers.GetFilmCategoriesHandler)
	server.GET("/film-category/:film_id/:category_id", handlers.GetFilmCategoryHandler)
	server.PUT("/film-category/:film_id/:category_id", handlers.PutFilmCategoryHandler)
	server.DELETE("/film-category/:film_id/:category_id", handlers.DeleteFilmCategoryHandler)
}
