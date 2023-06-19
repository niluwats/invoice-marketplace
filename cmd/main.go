package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/niluwats/invoice-marketplace/internal/handlers"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	handlers.StartServer()
}
