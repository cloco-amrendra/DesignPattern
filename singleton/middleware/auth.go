// singleton/middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is the Basic Authentication middleware
func AuthMiddleware(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()
		if !hasAuth || user != username || pass != password {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
