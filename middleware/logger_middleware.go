package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		duration := time.Since(start)
		log.Printf("Request %s %s took %v", context.Request.Method, context.Request.URL.Path, duration)
	}
}
