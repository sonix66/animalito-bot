package announcementservice

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
)

func (s *Service) GetAnnouncementByID(ctx context.Context, id string) (*entity.Announcement, error) {
	announcement, err := s.repo.GetAnnouncementByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.repo.GetAnnouncementByID: %w", err)
	}

	return announcement, nil
}
