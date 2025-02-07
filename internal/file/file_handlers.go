package file

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetFileHandler(context *gin.Context) {
	filepath := context.Param("filepath")

	err := isValidFilePath(filepath)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileFound, err := isFileExists(currentDir + filepath)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"path": filepath, "directory": currentDir, "found": fileFound})
}
