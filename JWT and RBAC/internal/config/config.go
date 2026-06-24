package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI  string `env:"MONGO_URI"`
	MongoDB   string `env:"MONGO_DB_NAME"`
	Port      string `env:"PORT"`
	JWTSecret string `env:"JWT_SECRET"`
}

func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		MongoURI:  strings.TrimSpace(os.Getenv("MONGO_URI")),
		MongoDB:   strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		Port:      strings.TrimSpace(os.Getenv("PORT")),
		JWTSecret: strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}
	if cfg.MongoURI == "" || cfg.MongoDB == "" || cfg.Port == "" || cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("missing required environment variables")
	}
	return cfg, nil
}
