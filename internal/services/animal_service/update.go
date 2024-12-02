package animalservice

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (s *Service) UpdateAnimalByID(ctx context.Context, animal *entity.Animal) error {
	err := s.repo.UpdateAnimal(ctx, animal)
	if err != nil {
		return fmt.Errorf("s.repo.UpdateAnimal: %w", err)
	}
	return nil
}
