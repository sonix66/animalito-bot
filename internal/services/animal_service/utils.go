package animalservice

import (
	"fmt"
	"os"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (s *Service) savePhoto(name string, data []byte) error {
	filePath := fmt.Sprintf("%s/%s", s.staticFolder, name)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("file.Write: %w", err)
	}
	return nil
}

func (s *Service) savePhotoPreparedData(data []byte) entity.OnSaveCallback {
	return func(fileName string) error {
		return s.savePhoto(fileName, data)
	}
}

func (s *Service) getStaticURL(fileName string) string {
	return fmt.Sprintf("%s/%s", s.photosPrefixURL, fileName)
}

func (s *Service) getFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", s.staticFolder, fileName)
}