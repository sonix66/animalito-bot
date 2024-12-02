package animalservice

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (s *Service) CreateAnimal(ctx context.Context, animal *entity.Animal, filesData [][]byte) error {
	animalId, err := s.repo.CreateAnimal(ctx, animal)
	if err != nil {
		return fmt.Errorf("s.repo.CreateAnimal: %w", err)
	}
	
	for _, fileData := range filesData{
		err = s.repo.CreatePhoto(ctx, animalId, s.savePhotoPreparedData(fileData))
		if err != nil {
			return fmt.Errorf("s.repo.CreatePhoto: %w", err)
		}
		
	}
	return nil
}
