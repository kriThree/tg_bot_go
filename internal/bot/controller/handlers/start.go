package handlers

import (
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *BotHandlers) Start(update tgbotapi.Update, state *statemanager.UserState) {

	if update.Message.Chat.IsSuperGroup() || update.Message.Chat.IsChannel() || update.Message.Chat.IsGroup() {
		h.api.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Unfortunately our bot can't write to the group \n please go to private messages",
		))
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"Have a nice day "+update.Message.From.FirstName+"\n"+"Хотите заниматься английским? Я могу вам с этим помочь!")
	btn1 := tgbotapi.NewKeyboardButton(SAVE_DEFINITION_QUERY)
	btn2 := tgbotapi.NewKeyboardButton("Get definitions")

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(btn1, btn2))

	state.Operation = statemanager.BASE
	_, err := h.api.Send(msg)

	if err != nil {
		h.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Error"))
	}
}
