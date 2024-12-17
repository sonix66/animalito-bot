package tgbot

import (
	"context"
	"errors"
	"fmt"

	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/sonix66/animalito-bot/internal/entity"
)

func (c *Controller) SwitchAnimalCallbackHandler(
	ctx context.Context,
	cbq *tgb.CallbackQueryUpdate,
	cbd NextAnimalCallbackData,
) error {
	totalCount, err := c.AnimalService.GetAnimalsCount(ctx)
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
		return errors.New("animal list is empty")
	}
	messageText := fmt.Sprintf(
		"%s\n\n%s\n\nОбъявление от: %s",
		animals[0].Name,
		animals[0].Description,
		animals[0].CreatedAt.Local().Format("02.01.2006"),
	)
	keyboard := tg.NewInlineKeyboardMarkup(
		[]tg.InlineKeyboardButton{
			nextAnimalDataFilter.MustButton(
				"⬅️ Назад",
				NextAnimalCallbackData{
					Current: cbd.Current - 1,
				},
			),
			nextAnimalDataFilter.MustButton(
				"➡️ Вперед",
				NextAnimalCallbackData{
					Current: cbd.Current + 1,
				},
			),
		},
		[]tg.InlineKeyboardButton{
			tg.NewInlineKeyboardButtonURL(
				"☎️",
				"https://nekrasovka-priut.ru/",
			),
			tg.NewInlineKeyboardButtonURL(
				"💰",
				"https://nekrasovka-priut.ru/",
			),
		},
	)
	c.mu.RLock()
	photoFileID, ok := c.localToTGPhotoIDMap[animals[0].PhotoURLs[0]]
	c.mu.RUnlock()
	if !ok {
		inputFile, errNewInputFileLocal := tg.NewInputFileLocal(animals[0].PhotoURLs[0])
		if errNewInputFileLocal != nil {
			return fmt.Errorf("tg.NewInputFileLocal: %w", errNewInputFileLocal)
		}
		defer inputFile.Close()
		sendPhotoCall := cbq.Client.SendPhoto(
			tg.ChatID(c.cfg.PreloadPhotoChatID),
			tg.NewFileArgUpload(inputFile),
		)
		photoMessage, errDo := sendPhotoCall.Do(ctx)
		if errDo != nil {
			return fmt.Errorf("sendPhotoCall.Do: %w", errDo)
		}
		if len(photoMessage.Photo) == 0 {
			return errors.New("no photo returned from Telegram")
		}

		photoFileID = string(photoMessage.Photo[0].FileID)

		c.mu.Lock()
		c.localToTGPhotoIDMap[animals[0].PhotoURLs[0]] = photoFileID
		c.mu.Unlock()
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
