package userUseCase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mensina-be/config"
	"mensina-be/database"
	"mensina-be/database/models"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type APIResponse struct {
	SecureURL string `json:"secure_url"`
}

func UploadUserIcon(userId uint, icon []byte) *config.RestErr {
	cloudName := os.Getenv("CLOUD_NAME")
	preset := os.Getenv("UPLOAD_PRESET")

	go deleteUserIcon(userId, cloudName)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("upload_preset", preset)

	// Criar o campo para a imagem
	part, err := writer.CreateFormFile("file", fmt.Sprintf("usericon_%d.png", userId))
	if err != nil {
		return config.NewInternaErr("Erro ao criar campo icon no FormData")
	}

	// Escrever a imagem no campo
	_, err = io.Copy(part, bytes.NewReader(icon))
	if err != nil {
		return config.NewInternaErr("Erro ao adicionar imagem ao FormData")
	}

	// Finalizar o writer para concluir o FormData
	err = writer.Close()
	if err != nil {
		return config.NewInternaErr("Erro ao finalizar o FormData")
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/image/upload", cloudName), body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return config.NewInternaErr("Erro ao enviar requisição para o servidor")
	}
	defer resp.Body.Close()
	fmt.Printf("%+v\n", resp.Body)

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return config.NewInternaErr("Erro ao processar resposta da API")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro da API: %s\n", string(resBody))
		return config.NewInternaErr("Erro ao enviar imagem: " + string(resBody))
	}

	// Decodificar o JSON da resposta
	var response APIResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return config.NewInternaErr("Erro ao decodificar JSON: " + err.Error())
	}

	db := database.GetDatabase()

	dbErr := db.
		Model(models.User{}).
		Where("id = ?", userId).
		Update("safe_url_icon", response.SecureURL).
		Error

	if dbErr != nil {
		return config.NewInternaErr("Erro ao salvar safeUrl no banco de dados")
	}

	return nil
}

func deleteUserIcon(userId uint, cloudName string) {
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	user, _ := GetUserInfos(userId)

	if user.SafeUrlIcon == "" {
		return
	}

	urlParts := strings.Split(user.SafeUrlIcon, "/")
	publicID := strings.TrimSuffix(urlParts[len(urlParts)-1], filepath.Ext(urlParts[len(urlParts)-1]))

	body := map[string]string{
		"public_ids": publicID,
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/resources/image/upload", cloudName), bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(apiKey, apiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
