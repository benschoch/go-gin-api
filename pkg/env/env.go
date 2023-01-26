package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	MongoDBURL string
}

func LoadEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No .env file found")
	}

	mongoDBURL := os.Getenv("MONGODB_URL")

	// TODO: add validation of environment variables
	return &Environment{
		MongoDBURL: mongoDBURL,
	}
}
