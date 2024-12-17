package http

import (
	"context"

	"github.com/sonix66/animalito-bot/internal/entity"
)

type (
	Config struct {
		Host            string `mapstructure:"host"`
		Port            string `mapstructure:"port"`
		StaticFolder    string `mapstructure:"static_folder"`
		PhotosPrefixURL string `mapstructure:"photos_prefix_url"`
	}

	AnimalService interface {
		CreateAnimal(ctx context.Context, animal *entity.Animal, filesData [][]byte) error
		GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error)
		GetAnimalList(
			ctx context.Context,
			count int,
			offset int,
			photoSource entity.PhotoSource,
		) ([]*entity.Animal, error)
		UpdateAnimalByID(ctx context.Context, animal *entity.Animal) error
		DeleteAnimalByID(ctx context.Context, id string) error
		GetAnimalsCount(ctx context.Context) (int, error)
	}

	Controller struct {
		animalService AnimalService
		cfg           *Config
	}
)

func New(animalService AnimalService, cfg *Config) *Controller {
	return &Controller{
		animalService: animalService,
		cfg:           cfg,
	}
}
