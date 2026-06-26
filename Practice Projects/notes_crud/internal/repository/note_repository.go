package repository

import (
	"context"
	"errors"
	"fmt"
	noteModel "notes-api/internal/models/note"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type NoteRepository struct {
	collection *mongo.Collection
}

func NewNoteRepository(db *mongo.Database) *NoteRepository {
	return &NoteRepository{
		collection: db.Collection("notes"),
	}
}

func (r *NoteRepository) Create(ctx context.Context, note noteModel.Note) (noteModel.Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(opCtx, note)
	if err != nil {
		return noteModel.Note{}, err
	}

	return note, nil
}

func (r *NoteRepository) GetAll(ctx context.Context) ([]noteModel.Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{}
	cursor, err := r.collection.Find(opCtx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve notes: %w", err)
	}
	defer cursor.Close(opCtx)

	var notes []noteModel.Note
	if err := cursor.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("failed to decode notes: %w", err)
	}
	return notes, nil
}

func (r *NoteRepository) GetByID(ctx context.Context, id bson.ObjectID) (noteModel.Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var note noteModel.Note
	err := r.collection.FindOne(opCtx, bson.M{"_id": id}, options.FindOne()).Decode(&note)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return noteModel.Note{}, fmt.Errorf("note not found")
		}
		return noteModel.Note{}, fmt.Errorf("failed to retrieve note: %w", err)
	}
	return note, nil
}

func (r *NoteRepository) Update(ctx context.Context, id bson.ObjectID, req noteModel.UpdateNoteRequest) (noteModel.Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateFields := bson.M{}
	if req.Title != nil {
		updateFields["title"] = *req.Title
	}
	if req.Content != nil {
		updateFields["content"] = *req.Content
	}
	if req.Pinned != nil {
		updateFields["pinned"] = *req.Pinned
	}
	updateFields["updated_at"] = time.Now().UTC()
	update := bson.M{
		"$set": updateFields,
	}

	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var updatedNote noteModel.Note

	err := r.collection.
		FindOneAndUpdate(opCtx, bson.M{"_id": id}, update, opts).
		Decode(&updatedNote)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return noteModel.Note{}, fmt.Errorf("note not found")
		}
		return noteModel.Note{}, fmt.Errorf("failed to update note: %w", err)
	}

	return updatedNote, nil
}

func (r *NoteRepository) Delete(ctx context.Context, id bson.ObjectID) (bool, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := r.collection.DeleteOne(opCtx, bson.M{"_id": id})
	if err != nil {
		return false, fmt.Errorf("failed to delete note: %w", err)
	}
	if res.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}
