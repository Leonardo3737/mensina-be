package userController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Obter ícone de usuário
// @Tags User
// @Security BearerAuth
// @Param user_id path int true "USER ID"
// @Produce  png
// @Success 200 {file} file "Ícone do usuário em PNG"
// @Router /user/icon/{user_id} [get]
func GetUserIconById(c *gin.Context) {
	_userId := c.Param("user_id")
	userId, err := strconv.Atoi(_userId)

	if err != nil {
		restErr := config.NewBadRequestErr("ID has to be integer")
		c.JSON(restErr.Code, restErr)
		return
	}

	// Busca o ícone do usuário
	iconFile, restErr := userUseCase.GetUserIconById(uint(userId))
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	// Define o cabeçalho de tipo de conteúdo para imagem PNG
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.WriteHeader(200)
	c.Writer.Write(iconFile)
}
