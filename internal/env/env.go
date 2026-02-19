package env

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func GetString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
