package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if os.Getenv("IS_DOCKER") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system env variables")
		}
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
