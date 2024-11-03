package main

import (
	"log"
	"mensina-be/database"
	"mensina-be/server"

	"github.com/joho/godotenv"
)

// @title API Mensina
// @version 1.0
// @description API desenvolvida para projeto academico
// @host localhost:5000
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartDb()
	server := server.NewServer()
	server.Run()
}
