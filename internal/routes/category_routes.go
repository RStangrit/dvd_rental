package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(server *gin.Engine) {
	server.POST("/category", handlers.PostCategoryHandler)
	server.GET("/categories", handlers.GetCategoriesHandler)
	server.GET("/category/:id", handlers.GetCategoryHandler)
	server.PUT("/category/:id", handlers.PutCategoryHandler)
	server.DELETE("/category/:id", handlers.DeleteCategoryHandler)
}
