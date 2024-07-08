package handlers

import (
	"english_learn/internal/bot/controller/utils"
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *BotHandlers) Start(ctx utils.AppContext) {

	if ctx.Update.Message.Chat.IsSuperGroup() || ctx.Update.Message.Chat.IsChannel() || ctx.Update.Message.Chat.IsGroup() {
		h.api.Send(tgbotapi.NewMessage(
			ctx.Update.Message.Chat.ID,
			"Unfortunately our bot can't write to the group \n please go to private messages",
		))
		return
	}

	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID,
		"Have a nice day "+ctx.Update.Message.From.FirstName+"\n"+"Хотите заниматься английским? Я могу вам с этим помочь!")
	btn1 := tgbotapi.NewKeyboardButton(SAVE_DEFINITION_QUERY)
	btn2 := tgbotapi.NewKeyboardButton("Get definitions")

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(btn1, btn2))

	ctx.State.Operation = statemanager.BASE
	_, err := h.api.Send(msg)

	if err != nil {
		h.api.Send(tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, "Error"))
	}
}
