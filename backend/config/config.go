package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func GetFrontendOrigin() string {
	origin := os.Getenv("FRONTEND_ORIGIN")
	if origin == "" {
		origin = "http://localhost:5173"
	}
	return origin
}