package main

import (
	"log"
	"mensina-be/core/routines"
	"mensina-be/database"
	"mensina-be/server"

	"github.com/joho/godotenv"
)

// @title API Mensina
// @version 1.0
// @description API desenvolvida para projeto academico
// @host localhost:5000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Insira o token de autenticação no formato: "Bearer {token}"
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartDb()
	server := server.NewServer()

	callbackChannel := make(chan routines.RoutineCallback)

	go routines.RunQuizRoutine(callbackChannel)

	server.Run(callbackChannel)
}
