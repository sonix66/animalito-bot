package telegram

import "gopkg.in/telebot.v4"

func InitHandlers(bot *telebot.Bot, controller *Controller) {
	bot.Handle("/start", controller.HandleStartCommand)
	bot.Handle("/app", controller.HandleGetWebappButton)
}
