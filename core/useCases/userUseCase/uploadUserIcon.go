package userUseCase

import (
	"fmt"
	"mensina-be/config"
	"mensina-be/database"
	"mensina-be/database/models"
	"os"
	"path/filepath"
	"time"

	"github.com/olahol/go-imageupload"
)

func UploadUserIcon(userId uint, icon *imageupload.Image) *config.RestErr {
	fileName := fmt.Sprintf("user_%d_%d.png", userId, time.Now().Unix())
	uploadDir := filepath.Join("uploads", "user_icons")
	filePath := filepath.Join(uploadDir, fileName)

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return config.NewInternaErr("Erro ao criar diretório de upload")
	}

	// Salvar a imagem como PNG no sistema de arquivos
	err = icon.Save(filePath)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return config.NewInternaErr("Erro ao salvar a imagem")
	}

	// Atualizar o caminho do ícone no banco de dados
	db := database.GetDatabase()
	err = db.Model(&models.User{}).Where("id = ?", userId).Update("icon_path", filePath).Error

	if err != nil {
		return config.NewInternaErr("Erro ao atualizar o caminho da imagem no banco de dados")
	}
	return nil
}
