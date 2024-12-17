package animalservice

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (s *Service) GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error) {
	animal, err := s.repo.GetAnimalByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.repo.GetAnimalByID: %w", err)
	}

	for i, photoName := range animal.PhotoURLs {
		animal.PhotoURLs[i] = s.getStaticURL(photoName)
	}

	return animal, nil
}
