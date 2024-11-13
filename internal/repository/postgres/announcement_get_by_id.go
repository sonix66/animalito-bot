package postgres

import (
	"context"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (r *Repository) GetAnnouncementByID(ctx context.Context, id string) (*entity.Announcement, error) {
	return &entity.Announcement{
		ID: id,
	}, nil
}
