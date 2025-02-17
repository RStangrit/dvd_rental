package middleware

import (
	"bytes"
	"io"
	"log"
	"main/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	zerolog := logger.GetZerologger()

	return func(context *gin.Context) {
		body, _ := context.GetRawData()
		log.Printf("Method: %s\nPath: %s\nHeaders: %v\nBody: %s\n\n",
			context.Request.Method, context.Request.URL.Path, context.Request.Header, string(body))
		file, err := os.OpenFile("requestsLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)

		zerolog.Info().Str("method", context.Request.Method).
			Str("path", context.Request.URL.Path).
			Interface("headers", context.Request.Header).
			Str("body", string(body)).
			Msg("HTTP Request")

		context.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		context.Next()
	}
}
