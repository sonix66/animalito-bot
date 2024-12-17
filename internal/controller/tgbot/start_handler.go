package tgbot

import (
	"context"
	"errors"
	"fmt"

	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/sonix66/animalito-bot/internal/entity"
)

func (c *Controller) StartHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	animals, err := c.AnimalService.GetAnimalList(ctx, 1, 0, entity.PhotoSourceFile)
	if err != nil {
		return fmt.Errorf("c.announcementService.GetAnnouncementByID: %w", err)
	}
	if len(animals) == 0 {
		return errors.New("animal list is empty")
	}

	// Кнопки для навигации
	keyboard := tg.NewInlineKeyboardMarkup(
		[]tg.InlineKeyboardButton{
			nextAnimalDataFilter.MustButton(
				"⬅️ Назад",
				NextAnimalCallbackData{
					Current: -1,
				},
			),
			nextAnimalDataFilter.MustButton(
				"➡️ Вперед",
				NextAnimalCallbackData{
					Current: 1,
				},
			),
		},
		[]tg.InlineKeyboardButton{
			tg.NewInlineKeyboardButtonURL(
				"☎️",
				// fmt.Sprintf("https://t.me/%s", c.cfg.AdminUsername),
				"https://nekrasovka-priut.ru/",
			),
			tg.NewInlineKeyboardButtonURL(
				"💰",
				"https://nekrasovka-priut.ru/",
			),
		},
	)

	photoFileID, ok := c.localToTGPhotoIDMap[animals[0].PhotoURLs[0]]
	if !ok {
		inputFile, errNewInputFileLocal := tg.NewInputFileLocal(animals[0].PhotoURLs[0])
		if errNewInputFileLocal != nil {
			return fmt.Errorf("tg.NewInputFileLocal: %w", errNewInputFileLocal)
		}
		defer inputFile.Close()

		// Отправляем фото в Telegram
		sendPhotoCall := msg.Client.SendPhoto(
			tg.ChatID(c.cfg.PreloadPhotoChatID),
			tg.NewFileArgUpload(inputFile),
		)

		photoMessage, errDo := sendPhotoCall.Do(ctx)
		if errDo != nil {
			return fmt.Errorf("sendPhotoCall.Do: %w", errDo)
		}

		// Извлекаем FileID из первой версии фото
		if len(photoMessage.Photo) == 0 {
			return errors.New("no photo returned from Telegram")
		}

		photoFileID = string(photoMessage.Photo[0].FileID)
		c.localToTGPhotoIDMap[animals[0].PhotoURLs[0]] = photoFileID
	}

	answer := msg.Client.SendPhoto(
		msg.Chat.ID,
		tg.NewFileArgID(tg.FileID(photoFileID)),
	).Caption(
		tg.HTML.Text(
			fmt.Sprintf(
				"%s\n\n%s\n\nОбъявление от: %s",
				animals[0].Name,
				animals[0].Description,
				animals[0].CreatedAt.Local().Format("02.01.2006"),
			)),
	).ReplyMarkup(keyboard)

	return answer.DoVoid(ctx)
}
