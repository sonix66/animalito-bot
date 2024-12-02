package telegram

import (
	"context"

	"github.com/sonix66/animalito-bot/internal/entity"
)

type (
	Config struct {
		Token                string `mapstructure:"token"`
		PollerTimeoutSeconds int    `mapstructure:"poller_timeout_seconds"`
		AdminUsername        string `mapstructure:"admin_username"`
	}
	AnimalService interface {
		GetAnimalList(
			ctx context.Context,
			count int,
			offset int,
			photoSource entity.PhotoSource,
		) ([]*entity.Animal, error)
		GetAnimalCount(ctx context.Context) (int, error)
	}

	Controller struct {
		AnimalService AnimalService
		cfg           *Config
	}
)

func New(cfg *Config, AnimalService AnimalService) *Controller {
	return &Controller{
		AnimalService: AnimalService,
		cfg:           cfg,
	}
}
