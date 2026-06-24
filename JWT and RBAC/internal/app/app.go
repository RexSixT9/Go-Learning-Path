package app

import (
	"context"
	"go-auth/internal/config"
	"go-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config config.Config

	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	mongoCli, err := db.ConnectMongo(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:      cfg,
		MongoClient: mongoCli.Client,
		DB:          mongoCli.DB,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient != nil {
		return a.MongoClient.Disconnect(ctx)
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return err
	}

	return nil
}
