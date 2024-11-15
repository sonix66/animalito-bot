package telegram

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/telebot.v4"
)

func (c *Controller) HandleStartCommand(ctx telebot.Context) error {
	goCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	announcement, err := c.announcementService.GetAnnouncementByID(goCtx, uuid.NewString())
	if err != nil {
		return fmt.Errorf("c.announcementService.GetAnnouncementByID: %w", err)
	}

	err = ctx.Reply(announcement.ID)
	if err != nil {
		return fmt.Errorf("ctx.Reply: %w", err)
	}

	return nil
}
