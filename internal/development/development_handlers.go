package development

import (
	apiclient "main/pkg/api_client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DevelopmentHandler struct {
	service *DevelopmentService
}

func NewDevelopmentHandler(service *DevelopmentService) *DevelopmentHandler {
	return &DevelopmentHandler{service: service}
}

func (handler *DevelopmentHandler) GetTestHandler(context *gin.Context) {
	apiclient.Heimdall(context)
	context.JSON(http.StatusOK, gin.H{"message": "Test request processed"})
}

func (handler *DevelopmentHandler) GetReindexFilmsHandler(context *gin.Context) {
	_, err := handler.service.ReadAllFilmsForIndexing()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "data successfully indexed"})
}
