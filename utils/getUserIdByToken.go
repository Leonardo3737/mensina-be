package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserIdByToken(c *gin.Context) (uint, error) {
	userId, exists := c.Get("userId")
	if !exists {
		return 0, fmt.Errorf("user id not found")
	}

	id, ok := userId.(uint)

	if !ok {
		return 0, fmt.Errorf("id has to be integer")
	}
	return id, nil
}
