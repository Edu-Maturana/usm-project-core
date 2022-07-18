package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	if os.Getenv("ENVIRONMENT") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}

		return os.Getenv(key)
	} else if os.Getenv("ENVIRONMENT") == "prod" {
		return os.Getenv(key)
	} else {
		panic("Invalid environment")
	}
}
