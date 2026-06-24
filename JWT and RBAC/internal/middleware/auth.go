package middleware

import (
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDkey = "auth.userId"
	ctxRolekey   = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(scheme, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			return
		}

		claims, err := auth.ParseToken(jwtSecret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		c.Set(ctxUserIDkey, claims.Subject)
		c.Set(ctxRolekey, claims.Role)

		c.Next()
	}
}

func GetUserID(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxUserIDkey)
	if !ok {
		return "", false
	}

	userID, ok := res.(string)
	if !ok {
		return "", false
	}
	return userID, ok
}

func GetUserRole(c *gin.Context) (string, bool) {
	role, ok := c.Get(ctxRolekey)
	if !ok {
		return "", false
	}

	roleStr, ok := role.(string)
	if !ok {
		return "", false
	}
	return roleStr, ok
}
