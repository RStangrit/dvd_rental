package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeTrackerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		duration := time.Since(start)
		fmt.Printf("Request %s %s took %v", context.Request.Method, context.Request.URL.Path, duration)
	}
}
