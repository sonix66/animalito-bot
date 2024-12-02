package telegram

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sonix66/animalito-bot/internal/entity"
	"gopkg.in/telebot.v4"
)

func (c *Controller) handleNextButton(ctx telebot.Context) error {
	goCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	currentAnimal := ctx.Data()

	currentAnimalInt, err := strconv.Atoi(currentAnimal)
	if err != nil {
		return fmt.Errorf("invalid animal number")
	}

	animals, err := c.AnimalService.GetAnimalList(goCtx, 1, currentAnimalInt, entity.PhotoSourceFile)
	if err != nil {
		return fmt.Errorf("c.AnimalService.GetAnimalList: %w", err)
	}

	animalCount, err := c.AnimalService.GetAnimalCount(goCtx)
	if err != nil {
		return fmt.Errorf("c.AnimalService.GetAnimalCount: %w", err)
	}

	if animalCount == 0 {
		return ctx.Send("Нет объявлений")
	}

	if currentAnimalInt >= animalCount {
		ctx.Set(nextAnimalContextKey, 0)
		return c.handleNextButton(ctx)
	}

	if currentAnimalInt < 0 {
		ctx.Set(nextAnimalContextKey, animalCount-1)
		return c.handleNextButton(ctx)
	}

	photos, err := c.prepareAnimalToMassage(animals[0])
	if err != nil {
		return fmt.Errorf("c.prepareAnimalToMassage: %w", err)
	}

	ctx.Edit(photos[0], c.prepareKeyboard(currentAnimalInt))

	return nil
}
