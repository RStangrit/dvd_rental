package development

import (
	apiclient "main/pkg/api_client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DevelopmentHandler struct{}

func NewDevelopmentHandler(service *DevelopmentService) *DevelopmentHandler {
	return &DevelopmentHandler{}
}

func (handler *DevelopmentHandler) GetTestHandler(context *gin.Context) {
	apiclient.Heimdall(context)
	context.JSON(http.StatusOK, gin.H{"message": "Test request processed"})
}
