package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NewsAPIKey string
	Port       string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	config := &Config{
		NewsAPIKey: os.Getenv("NEWS_API_KEY"),
		Port:       os.Getenv("PORT"),
	}

	if config.NewsAPIKey == "" {
		log.Fatal("NEWS_API_KEY is required")
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	return config
}