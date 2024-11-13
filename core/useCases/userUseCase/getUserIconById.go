package userUseCase

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetUserIconById(userId uint) []byte {
	fileName := fmt.Sprintf("usericon_%d.png", userId)
	uploadDir := filepath.Join("uploads", "user_icons")
	filePath := filepath.Join(uploadDir, fileName)

	iconFile, err := os.ReadFile(filePath)

	if err != nil {
		return nil
	}
	return iconFile
}
