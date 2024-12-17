package tgbot

import (
	"context"
	"sync"

	"github.com/sonix66/animalito-bot/internal/entity"
)

type (
	Config struct {
		Token                string `mapstructure:"token"`
		PollerTimeoutSeconds int    `mapstructure:"poller_timeout_seconds"`
		AdminUsername        string `mapstructure:"admin_username"`
		WebAppURL            string `mapstructure:"web_app_url"`
		PreloadPhotoChatID   int64  `mapstructure:"preload_photo_chat_id"`
	}

	AnimalService interface {
		GetAnimalList(
			ctx context.Context,
			count int,
			offset int,
			photoSource entity.PhotoSource,
		) ([]*entity.Animal, error)
		GetAnimalsCount(ctx context.Context) (int, error)
	}

	Controller struct {
		AnimalService       AnimalService
		cfg                 *Config
		localToTGPhotoIDMap map[string]string
		mu                  *sync.RWMutex
	}
)

func New(cfg *Config, AnimalService AnimalService) *Controller {
	return &Controller{
		AnimalService:       AnimalService,
		cfg:                 cfg,
		localToTGPhotoIDMap: map[string]string{},
		mu:                  &sync.RWMutex{},
	}
}
