package development

import (
	apiclient "main/pkg/api_client"

	"github.com/gin-gonic/gin"
)

func GetTestHandler(context *gin.Context) {
	apiclient.Heimdall(context)
}
