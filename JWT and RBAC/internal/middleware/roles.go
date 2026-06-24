package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := GetUserRole(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role information is missing"})
			return
		}

		if !strings.EqualFold(role, "admin") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
			return
		}

		c.Next()
	}
}
