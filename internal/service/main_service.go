package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Hemanth5603/resume-go-server/configs"
	"github.com/Hemanth5603/resume-go-server/internal/repository"
)

type MainService interface {
	ForwardResumeAndDescriptionToModel(file multipart.File, filename, description string) error
}

type MainServiceImpl struct {
	mainServiceRepo *repository.MainServiceRepository
	config          *configs.Config
}

func NewMainService(mainServiceRepo *repository.MainServiceRepository, config *configs.Config) MainService {
	return &MainServiceImpl{
		mainServiceRepo: mainServiceRepo,
		config:          config,
	}
}

// accepts file, description, and target URL from the server and forwards the request to the model
func (s *MainServiceImpl) ForwardResumeAndDescriptionToModel(file multipart.File, filename,
	description string) error {

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("description", description); err != nil {
		return err
	}

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	targetURL := s.config.ForwardURL

	req, err := http.NewRequest("POST", targetURL, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%v", string(respBody))
	}

	return nil
}
