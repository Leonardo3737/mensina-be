package routes

import (
	"mensina-be/controllers/loginController"
	"mensina-be/controllers/quizController"
	"mensina-be/controllers/rankController"
	"mensina-be/controllers/tagController"
	"mensina-be/controllers/userController"
	"mensina-be/core/routines"
	"mensina-be/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, quizRoutineChannel chan routines.RoutineCallback) *gin.Engine {

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
		user.GET("/icon/:user_id", middlewares.Auth(), userController.GetUserIconById)
	}

	// QUIZ routes
	quiz := router.Group("quiz", middlewares.Auth())
	{
		quiz.GET("/", quizController.GetQuiz)
		quiz.POST("/", quizController.CreateQuiz)
		quiz.GET("/questions/:quiz_id", quizController.GetQuestionByQuiz)
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
