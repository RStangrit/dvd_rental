package file

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	service *FileService
}

func NewFileHandler(service *FileService) *FileHandler {
	return &FileHandler{service: service}
}

func (handler *FileHandler) GetFileHandler(context *gin.Context) {
	filepath := context.Param("filepath")

	err := handler.service.IsValidFilePath(filepath)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileFound, err := handler.service.IsFileExists(currentDir + filepath)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"path": filepath, "directory": currentDir, "found": fileFound})
}
