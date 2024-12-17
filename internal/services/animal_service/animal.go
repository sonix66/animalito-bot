package animalservice

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

type (
	Repo interface {
		GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error)
		GetAnimalList(ctx context.Context, count, offset int) ([]*entity.Animal, error)
		CreateAnimal(ctx context.Context, animal *entity.Animal) (string, error)
		CreatePhoto(ctx context.Context, animalID string, callback entity.OnSaveCallback) error
		DeleteAnimalByID(ctx context.Context, animalID string) error
		UpdateAnimal(ctx context.Context, animal *entity.Animal) error
		GetAnimalsCount(ctx context.Context) (int, error)
	}

	Service struct {
		repo Repo

		photosPrefixURL string
		staticFolder    string
	}
)

func New(repo Repo, host, photosPrefixURL, staticFolder string) *Service {
	fullPhotosPrefixURL := fmt.Sprintf("%s%s", host, photosPrefixURL)
	return &Service{
		repo:            repo,
		photosPrefixURL: fullPhotosPrefixURL,
		staticFolder:    staticFolder,
	}
}
