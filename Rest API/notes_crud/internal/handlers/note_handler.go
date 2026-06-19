package handlers

import (
	"errors"
	"net/http"
	"notes-api/internal/models/note"
	"notes-api/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NoteHandler struct {
	repo *repository.NoteRepository
}

func NewNoteHandler(repo *repository.NoteRepository) *NoteHandler {
	return &NoteHandler{
		repo: repo,
	}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req note.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	now := time.Now().UTC()
	note := note.Note{
		ID:        bson.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}
	createdNote, err := h.repo.Create(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}
	c.JSON(http.StatusCreated, createdNote)
}

func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	notes, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *NoteHandler) GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	note, err := h.repo.GetByID(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var req note.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedNote, err := h.repo.Update(c.Request.Context(), objID, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, updatedNote)
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	deleted, err := h.repo.Delete(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
