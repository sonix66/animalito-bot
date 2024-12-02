package telegram

import (
	"context"
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
	"gopkg.in/telebot.v4"
)

func (c *Controller) HandleStartCommand(ctx telebot.Context) error {
	goCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	animals, err := c.AnimalService.GetAnimalList(goCtx, 1, 0, entity.PhotoSourceFile)
	if err != nil {
		return fmt.Errorf("c.announcementService.GetAnnouncementByID: %w", err)
	}
	if len(animals) == 0 {
		return fmt.Errorf("animal list is empty")
	}

	photos, err := c.prepareAnimalToMassage(animals[0])

	if err != nil {
		return fmt.Errorf("prepareAnnouncementToMassage: %w", err)
	}

	keyboard := c.prepareKeyboard(0)

	err = ctx.Send(photos[0], keyboard)
	// messages, err := ctx.Bot().SendAlbum(ctx.Sender(), photos)

	if err != nil {
		return fmt.Errorf("ctx.SendAlbum: %w", err)
	}


	return nil
}
