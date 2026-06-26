package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func LoadEnv() (Config, error) {

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %v", err)
	}

	mongoURI, err := extractEnvVariable("MONGO_URI")
	if err != nil {
		return Config{}, err
	}

	mongoDB, err := extractEnvVariable("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	serverPort, err := extractEnvVariable("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: serverPort,
	}, nil

}

func extractEnvVariable(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return value, nil
}
