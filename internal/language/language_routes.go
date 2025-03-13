package language

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LanguageRoutes struct {
	handler *LanguageHandler
}

func NewLanguageRoutes(db *gorm.DB) *LanguageRoutes {
	repo := NewLanguageRepository(db)
	service := NewLanguageService(repo)
	handler := NewLanguageHandler(service)

	return &LanguageRoutes{handler: handler}
}

func (route *LanguageRoutes) RegisterLanguageRoutes(server *gin.Engine) {
	server.POST("/language", route.handler.PostLanguageHandler)
	server.POST("/languages", route.handler.PostLanguagesHandler)
	server.GET("/languages", route.handler.GetLanguagesHandler)
	server.GET("/language/:id", route.handler.GetLanguageHandler)
	server.PUT("/language/:id", route.handler.PutLanguageHandler)
	server.DELETE("/language/:id", route.handler.DeleteLanguageHandler)
}
