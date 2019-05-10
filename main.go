package main

import (
	"github.com/bb3104/salmon/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadEnv()
	server.Init()
}

func loadEnv() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
