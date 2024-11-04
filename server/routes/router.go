package routes

import (
	"mensina-be/controllers/loginController"
	quizController "mensina-be/controllers/quizControllers"
	"mensina-be/controllers/userController"
	"mensina-be/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	// USER routes
	user := router.Group("user")
	{
		user.GET("/", userController.GetUsers)
		user.GET("/:id", userController.GetById)
		user.POST("/", userController.CreateUser)
		// Rotas protegidas
		user.PUT("/:id", middlewares.AuthById(), userController.UpdateUser)
		user.DELETE("/:id", middlewares.AuthById(), userController.DeleteUser)
	}

	// AUTH routes
	router.POST("login", loginController.Login)

	// QUIZ routes
	quiz := router.Group("quiz", middlewares.Auth())
	{
		quiz.GET("/", quizController.GetQuiz)
		quiz.GET("/questions/:quiz_id", quizController.GetQuestionByQuiz)
	}

	return router
}
