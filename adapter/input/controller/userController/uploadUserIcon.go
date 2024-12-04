package userController

import (
	"io"
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Upload de ícone de usuário
// @Tags User
// @Security BearerAuth
// @Accept  mpfd
// @Produce  png
// @Param file formData file true "Imagem do ícone do usuário (PNG, JPEG, etc.)"
// @Success 200 {string} string "Imagem processada com sucesso"
// @Router /user/icon [post]
func (controller userController) UploadUserIcon(c *gin.Context) {
	userId, err := utils.GetUserIdByToken(c)
	if err != nil {
		restErr := config.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	// Obter o arquivo do formulário
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		restErr := config.NewBadRequestErr("Erro ao processar o arquivo enviado")
		c.JSON(restErr.Code, restErr)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		restErr := config.NewInternaErr("Erro ao processar o arquivo enviado")
		c.JSON(restErr.Code, restErr)
		return
	}

	// Chamar o use case para processar o ícone do usuário
	userCaseErr := userUseCase.UploadUserIcon(userId, fileBytes)
	if userCaseErr != nil {
		c.JSON(userCaseErr.Code, userCaseErr)
		return
	}

	// Retornar sucesso
	c.Status(201)
}
