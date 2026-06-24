package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "Go Auth Service",
		"message": "Health check passed",
		"time":    time.Now().UTC(),
	})
}
