package server

import (
	"net/http"
	"notes-api/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(database *mongo.Database) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Service is running"})
	})

	routes.RegisterNoteRoutes(router, database)

	return router
}
