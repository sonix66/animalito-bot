package animalservice

import (
	"context"
	"fmt"
)

func (s *Service) GetAnimalCount(ctx context.Context) (int, error) {
	count, err := s.repo.GetAnimalCount(ctx)
	if err != nil {
		return 0, fmt.Errorf("s.repo.GetAnimalCount: %w", err)
	}

	return count, nil
}
