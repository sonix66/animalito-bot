package telegram

import (
	"gopkg.in/telebot.v4"
)

const (
	nextAnimalContextKey = "nextAnimal"
)

var (
	Keyboard   = &telebot.ReplyMarkup{}
	PrevButton = Keyboard.Data(
		"⬅️",
		"next",
		"2",
	)
)
