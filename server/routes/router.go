package routes

import (
	controllers "mensina-be/controllers/user"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	user := router.Group("user")
	{
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetById)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
	return router
}
