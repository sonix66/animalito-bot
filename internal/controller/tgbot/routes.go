package tgbot

import (
	"context"

	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/sonix66/animalito-bot/pkg/logger"
)

const (
	startCommand = "start"
	adminCommand = "admin"
)

func InitRoutes(router *tgb.Router, controller *Controller) {
	router.Error(func(_ context.Context, update *tgb.Update, err error) error {
		logger.Error("error when handling update",
			"update_id", update.ID,
			"error", err)
		return nil
	})

	router.
		Message(
			controller.StartHandler,
			tgb.Command(startCommand),
			tgb.ChatType(tg.ChatTypePrivate),
		).
		Message(
			controller.WebAppHandler,
			tgb.Command(adminCommand),
			tgb.ChatType(tg.ChatTypePrivate),
		).
		CallbackQuery(
			NextAnimalDataFilter.Handler(controller.SwitchAnimalCallbackHandler),
			NextAnimalDataFilter.Filter(),
		)
}
