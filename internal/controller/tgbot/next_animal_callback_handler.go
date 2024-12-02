package tgbot

import (
	"context"
	"fmt"

	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/sonix66/animalito-bot/internal/entity"
)

func (c *Controller) SwitchAnimalCallbackHandler(ctx context.Context, cbq *tgb.CallbackQueryUpdate, cbd NextAnimalCallbackData) error {
	totalCount, err := c.AnimalService.GetAnimalCount(ctx)
	if err != nil {
		return fmt.Errorf("c.AnimalService.GetAnimalCount: %w", err)
	}

	if cbd.Current < 0 {
		return c.SwitchAnimalCallbackHandler(ctx, cbq, NextAnimalCallbackData{
			Current: totalCount - 1,
		})
	}

	if cbd.Current >= totalCount {
		return c.SwitchAnimalCallbackHandler(ctx, cbq, NextAnimalCallbackData{
			Current: 0,
		})
	}

	animals, err := c.AnimalService.GetAnimalList(ctx, 1, cbd.Current, entity.PhotoSourceFile)
	if err != nil {
		return fmt.Errorf("c.announcementService.GetAnnouncementByID: %w", err)
	}
	if len(animals) == 0 {
		return fmt.Errorf("animal list is empty")
	}

	messageText := fmt.Sprintf(
		"%s\n\n%s\n\nОбъявление от: %s",
		animals[0].Name,
		animals[0].Description,
		animals[0].CreatedAt.Local().Format("02.01.2006 15:04:05"),
	)

	// Кнопки для навигации
	keyboard := tg.NewInlineKeyboardMarkup(
		[]tg.InlineKeyboardButton{
			NextAnimalDataFilter.MustButton(
				"⬅️ Назад",
				NextAnimalCallbackData{
					Current: cbd.Current - 1,
				},
			),
			tg.NewInlineKeyboardButtonURL(
				"☎️",
				fmt.Sprintf("https://t.me/%s", c.cfg.AdminUsername),
			),
			NextAnimalDataFilter.MustButton(
				"➡️ Вперед",
				NextAnimalCallbackData{
					Current: cbd.Current + 1,
				},
			),
		},
	)

	photoFileID, ok := c.localToTGPhotoIDMap[animals[0].PhotoURLs[0]]
	if !ok {
		inputFile, err := tg.NewInputFileLocal(animals[0].PhotoURLs[0])
		if err != nil {
			return err
		}
		defer inputFile.Close()

		// Отправляем фото в Telegram
		sendPhotoCall := cbq.Client.SendPhoto(
			tg.ChatID(c.cfg.PreloadPhotoChatID),
			tg.NewFileArgUpload(inputFile),
		)

		photoMessage, err := sendPhotoCall.Do(ctx)
		if err != nil {
			return fmt.Errorf("sendPhotoCall.Do: %w", err)
		}

		// Извлекаем FileID из первой версии фото
		if len(photoMessage.Photo) == 0 {
			return fmt.Errorf("no photo returned from Telegram")
		}

		photoFileID = string(photoMessage.Photo[0].FileID)
		c.localToTGPhotoIDMap[animals[0].PhotoURLs[0]] = photoFileID
	}

	media := &tg.InputMediaPhoto{
		Media:     tg.NewFileArgID(tg.FileID(photoFileID)),
		Caption:   messageText,
		ParseMode: tg.HTML,
	}

	msg := cbq.Message.Message
	editMessageMediaCall := cbq.Client.
		EditMessageMedia(media).
		// EditMessageCaption(msg.Chat.ID, msg.ID, messageText).
		MessageID(msg.ID).ChatID(msg.Chat.ID)

	return editMessageMediaCall.
		ReplyMarkup(keyboard).
		DoVoid(ctx)
}
