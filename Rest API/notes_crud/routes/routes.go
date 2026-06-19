package routes

import (
	"notes-api/internal/handlers"
	"notes-api/internal/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterNoteRoutes(router *gin.Engine, db *mongo.Database) {
	repo := repository.NewNoteRepository(db)
	noteHandler := handlers.NewNoteHandler(repo)

	noteRoutes := router.Group("/notes")
	{
		noteRoutes.POST("/", noteHandler.CreateNote)
		noteRoutes.GET("/", noteHandler.GetAllNotes)
		noteRoutes.GET("/:id", noteHandler.GetNoteByID)
		noteRoutes.PUT("/:id", noteHandler.UpdateNote)
		noteRoutes.DELETE("/:id", noteHandler.DeleteNote)
	}

}
