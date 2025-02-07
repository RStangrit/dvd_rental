package film_category

import (
	"github.com/gin-gonic/gin"
)

func RegisterFilmCategoryRoutes(server *gin.Engine) {
	server.POST("/film-category", PostFilmCategoryHandler)
	server.GET("/film-categories", GetFilmCategoriesHandler)
	server.GET("/film-category/:film_id/:category_id", GetFilmCategoryHandler)
	server.PUT("/film-category/:film_id/:category_id", PutFilmCategoryHandler)
	server.DELETE("/film-category/:film_id/:category_id", DeleteFilmCategoryHandler)
}
