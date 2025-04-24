package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// load .env files
func LoadConfig(filepaths ...string) {
	if err := godotenv.Load(filepaths...); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}

func GetAPIKey() string {
	return os.Getenv("DASHSCOPE_API_KEY")
}
