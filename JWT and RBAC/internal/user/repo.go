package user

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	col *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		col: db.Collection("users"),
	}
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	email = strings.ToLower(email)

	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, fmt.Errorf("user not found")
		}

		return User{}, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}

func (r *Repo) Create(ctx context.Context, user User) (User, error) {
	res, err := r.col.InsertOne(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}

	id, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	user.ID = id

	return user, nil
}
