package tgbot

import "github.com/mr-linch/go-tg/tgb"

var (
	//nolint:gochecknoglobals // not global
	nextAnimalDataFilter = tgb.NewCallbackDataFilter[NextAnimalCallbackData](
		"next_animal",
	)
)

type NextAnimalCallbackData struct {
	Current int
}
