package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadMongoENV() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: No .env file found")
	}
	return os.Getenv("MONGODB")
}
