package routes

import (
	"mensina-be/adapter/input/controller/loginController"
	"mensina-be/adapter/input/controller/quizController"
	"mensina-be/adapter/input/controller/rankController"
	"mensina-be/adapter/input/controller/tagController"
	"mensina-be/adapter/input/controller/userController"
	"mensina-be/adapter/input/server/middlewares"
	"mensina-be/core/routines"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(
	router *gin.Engine,
	quizRoutineChannel chan routines.RoutineCallback,
	userController userController.IUserController,
) *gin.Engine {

	// AUTH routes
	router.POST("login", loginController.Login)

	// VALIDATION TOKEN
	router.GET("/validate_token", middlewares.Auth())

	// USER routes
	user := router.Group("user")
	{
		user.GET("/", userController.GetUsers)
		user.POST("/", userController.CreateUser)
		// Rotas protegidas
		user.GET("/user_infos", middlewares.Auth(), userController.GetUserInfos)
		user.GET("/kpi", middlewares.Auth(), userController.GetUserKpi)
		user.PUT("/", middlewares.Auth(), userController.UpdateUser)
		user.DELETE("/", middlewares.Auth(), userController.DeleteUser)
		user.POST("/icon", middlewares.Auth(), userController.UploadUserIcon)
	}

	// QUIZ routes
	quiz := router.Group("quiz", middlewares.Auth())
	{
		quiz.GET("/", func(c *gin.Context) { quizController.GetQuiz(c, quizRoutineChannel) })
		quiz.POST("/", quizController.CreateQuiz)
		quiz.GET("/questions/:quiz_id", quizController.GetQuestionByQuiz)
		quiz.GET("/history", quizController.GetHistory)
		quiz.GET("/answer_check", func(c *gin.Context) { quizController.AnswerCheck(c, quizRoutineChannel) })
		quiz.GET("/start/:quiz_id", func(c *gin.Context) { quizController.StartQuiz(c, quizRoutineChannel) })
		quiz.DELETE("/finish/:quiz_id", func(c *gin.Context) { quizController.FinishQuiz(c, quizRoutineChannel) })
	}

	// TAG routes
	tag := router.Group("tag", middlewares.Auth())
	{
		tag.GET("/", tagController.GetTags)
		tag.POST("/", tagController.CreateTag)
	}

	// RANK routes
	rank := router.Group("rank", middlewares.Auth())
	{
		rank.GET("/", rankController.GetRank)
	}

	return router
}