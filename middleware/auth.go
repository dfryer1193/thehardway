package middleware

import (
	"github.com/dfryer1193/thehardway/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware is an example middleware that checks authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Validate the token
		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Continue to the next handler
		c.Next()
	}
}
