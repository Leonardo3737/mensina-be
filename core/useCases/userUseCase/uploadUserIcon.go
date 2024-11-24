package userUseCase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mensina-be/config"
	"mime/multipart"
	"net/http"
	"os"
)

type APIResponse struct {
	AssetID      string    `json:"asset_id"`
	PublicID     string    `json:"public_id"`
	Version      interface{}     `json:"version"`
	VersionID    string    `json:"version_id"`
	Signature    string    `json:"signature"`
	Width        interface{}       `json:"width"`
	Height       interface{}       `json:"height"`
	Format       string    `json:"format"`
	ResourceType string    `json:"resource_type"`
	CreatedAt    string    `json:"created_at"`
	Tags         []interface{}  `json:"tags"`
	Bytes        interface{}       `json:"bytes"`
	Type         string    `json:"type"`
	ETag         string    `json:"etag"`
	Placeholder  bool      `json:"placeholder"`
	URL          string    `json:"url"`
	SecureURL    string    `json:"secure_url"`
	AssetFolder  string    `json:"asset_folder"`
	DisplayName  string    `json:"display_name"`
}

func UploadUserIcon(userId uint, icon []byte) *config.RestErr {

	url := os.Getenv("URL_UPLOAD_IMAGE")
	preset := os.Getenv("UPLOAD_PRESET")

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

	req, _ := http.NewRequest("POST", url, body)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			return config.NewInternaErr("Erro ao enviar requisição para o servidor")
	}
	defer resp.Body.Close()

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

	fmt.Println(response.SecureURL)


	return nil
}
