package animalservice

import (
	"context"
	"fmt"
)

func (s *Service) DeleteAnimalByID(ctx context.Context, id string) error {
	err := s.repo.DeleteAnimalByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.repo.DeleteAnimal: %w", err)
	}

	return nil
}
