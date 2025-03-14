package film_category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FilmCategoryRoutes struct {
	handler *FilmCategoryHandler
}

func NewFilmCategoryRoutes(db *gorm.DB) *FilmCategoryRoutes {
	repo := NewFilmCategoryRepository(db)
	service := NewFilmCategoryService(repo)
	handler := NewFilmCategoryHandler(service)

	return &FilmCategoryRoutes{handler: handler}
}

func (route *FilmCategoryRoutes) RegisterFilmCategoryRoutes(server *gin.Engine) {
	server.POST("/film-category", route.handler.PostFilmCategoryHandler)
	server.GET("/film-categories", route.handler.GetFilmCategoriesHandler)
	server.GET("/film-category/:film_id/:category_id", route.handler.GetFilmCategoryHandler)
	server.PUT("/film-category/:film_id/:category_id", route.handler.PutFilmCategoryHandler)
	server.DELETE("/film-category/:film_id/:category_id", route.handler.DeleteFilmCategoryHandler)
}
