package routes

import (
	"mensina-be/controllers/loginController"
	"mensina-be/controllers/quizController"
	"mensina-be/controllers/tagController"
	"mensina-be/controllers/userController"
	"mensina-be/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	// AUTH routes
	router.POST("login", loginController.Login)

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

	// QUIZ routes
	quiz := router.Group("quiz", middlewares.Auth())
	{
		quiz.GET("/", quizController.GetQuiz)
		quiz.GET("/questions/:quiz_id", quizController.GetQuestionByQuiz)
		quiz.GET("/answer_check", quizController.AnswerCheck)
	}

	tag := router.Group("tag", middlewares.Auth())
	{
		tag.GET("/", tagController.GetTags)
	}

	return router
}
