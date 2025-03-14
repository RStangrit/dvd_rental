package file

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FileRoutes struct {
	handler *FileHandler
}

func NewFileRoutes(db *gorm.DB) *FileRoutes {
	service := NewFileService()
	handler := NewFileHandler(service)
	return &FileRoutes{handler: handler}
}

func (route *FileRoutes) RegisterFileRoutes(server *gin.Engine) {
	server.GET("/files/*filepath", route.handler.GetFileHandler)
}
