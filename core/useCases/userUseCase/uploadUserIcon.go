package userUseCase

import (
	"fmt"
	"mensina-be/config"
	"os"
	"path/filepath"

	"github.com/olahol/go-imageupload"
)

func UploadUserIcon(userId uint, icon *imageupload.Image) *config.RestErr {
	fileName := fmt.Sprintf("usericon_%d.png", userId)
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

	return nil
}
