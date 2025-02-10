package middleware

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		body, _ := context.GetRawData()
		fmt.Printf("Method: %s\nPath: %s\nHeaders: %v\nBody: %s\n",
			context.Request.Method, context.Request.URL.Path, context.Request.Header, string(body))

		context.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		context.Next()
	}
}
