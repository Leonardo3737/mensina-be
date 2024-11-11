package userController

import (
	"mensina-be/config"
	"mensina-be/core/useCases/userUseCase"
	"mensina-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

// @Summary Upload de ícone de usuário
// @Tags User
// @Security BearerAuth
// @Accept  mpfd
// @Produce  png
// @Param file formData file true "Imagem do ícone do usuário (PNG, JPEG, etc.)"
// @Success 200 {string} string "Imagem processada com sucesso"
// @Router /user/icon [post]
func UploadUserIcon(c *gin.Context) {
	userId, err := utils.GetUserIdByToken(c)
	if err != nil {
		restErr := config.NewBadRequestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	userIcon, err := imageupload.Process(c.Request, "file")

	if err != nil {
		restErr := config.NewInternaErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	userCaseErr := userUseCase.UploadUserIcon(userId, userIcon)

	if userCaseErr != nil {
		c.JSON(userCaseErr.Code, userCaseErr)
		return
	}

	userIconPng, err := imageupload.ThumbnailPNG(userIcon, 300, 300)

	if err != nil {
		restErr := config.NewInternaErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	userIconPng.Write(c.Writer)
}
