package main

import (
	"log"
	"mensina-be/adapter/input/controller/userController"
	"mensina-be/adapter/input/server/routes"
	"mensina-be/core/routines"
	"mensina-be/database"
	"mensina-be/docs"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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

type Server struct {
	port   string
	server *gin.Engine
}

// @title API Mensina
// @version 1.0
// @description API desenvolvida para projeto academico
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Insira o token de autenticação no formato: "Bearer {token}"
func main() {
	database := database.StartDb()

	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8080" // Valor padrão para desenvolvimento
	}
	docs.SwaggerInfo.Host = swaggerHost

	server := NewServer()

	callbackChannel := make(chan routines.RoutineCallback)

	go routines.RunQuizRoutine(callbackChannel)

	server.Run(callbackChannel, database)
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run(
	quizRoutineChannel chan routines.RoutineCallback,
	database *gorm.DB,
) {

	s.server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	_userController := initDependdecies(database)

	router := routes.ConfigRoutes(
		s.server,
		quizRoutineChannel,
		_userController,
	)

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Print("server is running: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}

func initDependdecies(
	database *gorm.DB,
) userController.IUserController {
	return userController.NewUserController()
}
