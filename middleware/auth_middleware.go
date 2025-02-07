package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		// if token == "" || !isValidToken(token) {
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
			context.Abort()
			return
		}

		// userRole := getUserRole(token)

		// if userRole != requiredRole {
		// 	context.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		// 	context.Abort()
		// 	return
		// }

		context.Next()
	}
}
