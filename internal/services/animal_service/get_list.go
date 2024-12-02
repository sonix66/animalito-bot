package animalservice

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (s *Service) GetAnimalList(
	ctx context.Context,
	count,
	offset int,
	photoSource entity.PhotoSource,
) ([]*entity.Animal, error) {
	animals, err := s.repo.GetAnimalList(ctx, count, offset)
	if err != nil {
		return nil, fmt.Errorf("s.repo.GetAnimalList: %w", err)
	}

	photoSourceFunc := s.getStaticURL
	if photoSource == entity.PhotoSourceFile {
		photoSourceFunc = s.getFilePath
	}

	for i, animal := range animals {
		for j, photoName := range animal.PhotoURLs {
			animals[i].PhotoURLs[j] = photoSourceFunc(photoName)
		}
	}

	return animals, nil
}
