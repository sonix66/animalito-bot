package animalservice

import (
	"context"
	"fmt"
)

func (s *Service) GetAnimalsCount(ctx context.Context) (int, error) {
	count, err := s.repo.GetAnimalsCount(ctx)
	if err != nil {
		return 0, fmt.Errorf("s.repo.GetAnimalCount: %w", err)
	}

	return count, nil
}
