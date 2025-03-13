package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRoutes struct {
	handler *CategoryHandler
}

func NewCategoryRoutes(db *gorm.DB) *CategoryRoutes {
	repo := NewCategoryRepository(db)
	service := NewCategoryService(repo)
	handler := NewCategoryHandler(service)

	return &CategoryRoutes{handler: handler}
}

func (route *CategoryRoutes) RegisterCategoryRoutes(server *gin.Engine) {
	server.POST("/category", route.handler.PostCategoryHandler)
	server.GET("/categories", route.handler.GetCategoriesHandler)
	server.GET("/category/:id", route.handler.GetCategoryHandler)
	server.PUT("/category/:id", route.handler.PutCategoryHandler)
	server.DELETE("/category/:id", route.handler.DeleteCategoryHandler)
}
