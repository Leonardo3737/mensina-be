package controllers

import (
	"mensina-be/core/models"
	"mensina-be/core/useCases/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := user.GetUsers()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot list users",
		})
		return
	}

	c.JSON(200, users)
}

func GetById(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	user, err := user.GetUserById(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find user",
		})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var _user models.User

	err := c.ShouldBindJSON(&_user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON",
		})
		return
	}

	newUser, status, err := user.CreateUser(&_user)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, newUser)
}

func UpdateUser(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	var _user models.User

	err = c.ShouldBindJSON(&_user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON",
		})
		return
	}

	if int(_user.ID) != id {
		c.JSON(400, gin.H{
			"error": "request ID must be equal to the body object ID",
		})
		return
	}

	updatedUser, status, err := user.UpdateUser(&_user)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, updatedUser)
}

func DeleteUser(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	err = user.DeleteUser(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find user",
		})
		return
	}

	c.Status(204)
}
