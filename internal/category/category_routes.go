package category

import (
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(server *gin.Engine) {
	server.POST("/category", PostCategoryHandler)
	server.GET("/categories", GetCategoriesHandler)
	server.GET("/category/:id", GetCategoryHandler)
	server.PUT("/category/:id", PutCategoryHandler)
	server.DELETE("/category/:id", DeleteCategoryHandler)
}
