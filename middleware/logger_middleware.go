package middleware

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		body, _ := context.GetRawData()
		log.Printf("Method: %s\nPath: %s\nHeaders: %v\nBody: %s\n\n",
			context.Request.Method, context.Request.URL.Path, context.Request.Header, string(body))
		file, err := os.OpenFile("requestsLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)

		context.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		context.Next()
	}
}
