package main

import (
	"log"
	"mensina-be/core/routines"
	"mensina-be/database"
	"mensina-be/server"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Carrega o .env somente em desenvolvimento
	if os.Getenv("DB_CONNECTION") == "" { // Use uma variável específica para identificar o ambiente
			err := godotenv.Load()
			if err != nil {
					log.Println("No .env file found, using Render environment variables")
			}
	}
}

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
	database.StartDb()
	server := server.NewServer()

	callbackChannel := make(chan routines.RoutineCallback)

	go routines.RunQuizRoutine(callbackChannel)

	server.Run(callbackChannel)
}
