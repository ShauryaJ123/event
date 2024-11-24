package middlewares

import (
	"net/http"
	"strings"

	"abc.com/calc/utils"
	"github.com/gin-gonic/gin"
)


func Authenticate(context *gin.Context) {
	// Get the Authorization header
	authHeader := context.Request.Header.Get("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	// Split the header to check for "Bearer" prefix
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}

	// Extract the token
	token := parts[1]

	// Verify the token
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
		return
	}

	// Store the user ID in the context
	context.Set("userId", userId)
	context.Next()
}
