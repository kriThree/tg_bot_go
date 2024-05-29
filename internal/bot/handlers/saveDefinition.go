package handlers

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *BotHandlers) SaveDefinition(update tgbotapi.Update) {

	req := update.Message.Text

	a := strings.Split(req, "-")

	if len(a) != 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong format")
		b.api.Send(msg)
		return
	}
	
	word := strings.Trim(a[0], " ")

	meaning := strings.Trim(a[1], " ")

	b.l.Info("Word: " + word + " Meaning: " + meaning)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.From.FirstName+" Hello")

	b.api.Send(msg)
}
