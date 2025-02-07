package middleware

import (
	"main/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")

		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
			context.Abort()
			return
		}

		if err := auth.VerifyToken(token); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Next()
	}
}
