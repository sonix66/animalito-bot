package telegram

import (
	"fmt"
	"net/url"

	"gopkg.in/telebot.v4"
)

func (c *Controller) HandleGetWebappButton(ctx telebot.Context) error {
	parsedURL, err := url.Parse(ctx.Message().Payload)

	if err != nil {
		ctx.Send("Плохой URL")
		return err
	}

	if parsedURL.Scheme != "https" {
		ctx.Send("Надо использовать https://")
		return fmt.Errorf("need to use https://")
	}

	var (
		webappKeyboard = &telebot.ReplyMarkup{}

		getWebappButton = webappKeyboard.WebApp(
			"Get Webapp",
			&telebot.WebApp{URL: parsedURL.String()},
		)
	)

	webappKeyboard.Inline(webappKeyboard.Row(getWebappButton))

	return ctx.Send("Заходи)", webappKeyboard)

}
