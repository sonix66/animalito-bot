package announcementservice

import (
	"context"

	"github.com/sonix66/animalito-bot/internal/entity"
)

type (
	Repo interface {
		GetAnnouncementByID(ctx context.Context, id string) (*entity.Announcement, error)
	}

	Service struct {
		repo Repo
	}
)

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}
