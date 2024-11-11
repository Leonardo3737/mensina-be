package userUseCase

import (
	"fmt"
	"mensina-be/config"
	"os"
	"path/filepath"
)

func GetUserIconById(userId uint) ([]byte, *config.RestErr) {
	fileName := fmt.Sprintf("usericon_%d.png", userId)
	uploadDir := filepath.Join("uploads", "user_icons")
	filePath := filepath.Join(uploadDir, fileName)

	iconFile, err := os.ReadFile(filePath)

	if err != nil {
		return nil, config.NewBadRequestErr(err.Error())
	}
	return iconFile, nil
}
