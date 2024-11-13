package telegram

import (
	"context"

	"github.com/sonix66/animalito-bot/internal/entity"
)

type (
	AnnouncementService interface {
		GetAnnouncementByID(ctx context.Context, id string) (*entity.Announcement, error)
	}

	Controller struct {
		announcementService AnnouncementService
	}
)

func New(announcementService AnnouncementService) *Controller {
	return &Controller{
		announcementService: announcementService,
	}
}
