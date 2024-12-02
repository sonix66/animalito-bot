package telegram

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v4"
)

func PrepareCurrentMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Callback() != nil && strings.HasPrefix(c.Callback().Data, "next____") {
			nextAnimal := strings.Split(c.Callback().Data, "____")
			if len(nextAnimal) != 2 {
				return fmt.Errorf("invalid callback data: %s", c.Callback().Data)
			}
			c.Set(nextAnimalContextKey, nextAnimal[1])
			c.Callback().Data = "next"
			c.Callback().Unique = "next"
		}
		return next(c) // continue execution chain
	}
}
