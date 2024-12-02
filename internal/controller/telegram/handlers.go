package telegram

import (
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func InitHandlers(bot *telebot.Bot, controller *Controller) {
	// bot.Use(PrepareCurrentMiddleware)
	bot.Use(middleware.Logger())
	bot.Handle(&telebot.InlineButton{Unique: "next"}, controller.handleNextButton)
	bot.Handle("/start", controller.HandleStartCommand)
	bot.Handle("/app", controller.HandleGetWebappButton)
}
