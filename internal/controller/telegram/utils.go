package telegram

import (
	"fmt"

	"github.com/sonix66/animalito-bot/internal/entity"
	"gopkg.in/telebot.v4"
)

func (c *Controller) prepareAnimalToMassage(animal *entity.Animal) (telebot.Album, error) {
	if len(animal.PhotoURLs) == 0 {
		return nil, ErrNoPhotos
	}
	photos := make(telebot.Album, 0, len(animal.PhotoURLs))
	for i, photoPath := range animal.PhotoURLs {
		telePhoto := &telebot.Photo{File: telebot.FromDisk(photoPath)}
		if i == 0 {
			telePhoto.Caption = fmt.Sprintf("*%s*\n%s\n%s\n\nОбъявление от: %s", animal.Name, animal.Type, animal.Description, animal.CreatedAt.Format("23.12.2024"))
		}
		photos = append(photos, telePhoto)
	}
	return photos, nil
}

func (c *Controller) prepareKeyboard(current int) *telebot.ReplyMarkup {
	// var (
	// 	keyboard = &telebot.ReplyMarkup{}

	// 	getPrevButton = keyboard.Data(
	// 		"⬅️",
	// 		"next",
	// 		fmt.Sprintf("%d", current-1),
	// 	)

	// 	getNextButton = keyboard.Data(
	// 		"➡️",
	// 		"next",
	// 		fmt.Sprintf("%d", current+1),
	// 	)

	// 	getContact = keyboard.URL("Написать админу", fmt.Sprintf("https://t.me/%s", c.cfg.AdminUsername))
	// )

	// keyboard.Inline(keyboard.Row(getPrevButton, getContact, getNextButton))
	Keyboard.Inline(Keyboard.Row(PrevButton, PrevButton, PrevButton))
	return Keyboard
}
