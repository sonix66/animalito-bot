package tgbot

import "github.com/mr-linch/go-tg/tgb"

var (
	NextAnimalDataFilter = tgb.NewCallbackDataFilter[NextAnimalCallbackData](
		"next_animal",
	)
)

type NextAnimalCallbackData struct {
	Current int
}
