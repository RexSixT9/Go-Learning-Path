package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Error closing app: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)
	srv := &http.Server{
		Addr:        ":5000",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("Starting server on port 5000...")

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
		}
		log.Fatalf("Error starting server: %v", err)
	}
}
