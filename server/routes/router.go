package routes

import (
	loginControllers "mensina-be/controllers/login"
	userControllers "mensina-be/controllers/user"
	"mensina-be/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	// USER routes
	user := router.Group("user")
	{
		user.GET("/", userControllers.GetUsers)
		user.GET("/:id", userControllers.GetById)
		user.POST("/", userControllers.CreateUser)
		// Rotas protegidas
		user.PUT("/:id", middlewares.AuthById(), userControllers.UpdateUser)
		user.DELETE("/:id", middlewares.AuthById(), userControllers.DeleteUser)
	}

	// AUTH routes
	router.POST("login", loginControllers.Login)

	return router
}
