package tgbot

import (
	"context"

	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func (c *Controller) WebAppHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	keyboard := tg.NewInlineKeyboardMarkup(
		[]tg.InlineKeyboardButton{
			{
				Text: "🖥️",
				WebApp: &tg.WebAppInfo{
					URL: c.cfg.WebAppURL,
				},
			},
		},
	)

	answer := msg.Client.SendMessage(
		msg.Chat.ID,
		"Нажмите кнопку ниже, чтобы перейти в админ-панель",
	).ReplyMarkup(keyboard)

	return answer.DoVoid(ctx)
}
