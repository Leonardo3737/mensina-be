package userController

import "github.com/gin-gonic/gin"

type IUserController interface {
	GetUserInfos(c *gin.Context)
	UploadUserIcon(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserKpi(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userController struct {
	service interface{}
}

func NewUserController() IUserController {
	return &userController{}
}
