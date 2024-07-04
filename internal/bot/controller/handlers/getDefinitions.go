package handlers

import (
	"context"
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *BotHandlers) GetDefinitions(update tgbotapi.Update, state *statemanager.UserState) {

	_, err := h.api.AnswerCallbackQuery(tgbotapi.CallbackConfig{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            "Definitions:",
	})

	h.hu.GetDefinitions(context.TODO())

	if err != nil {
		h.api.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Error"))
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Write your definition")

	state.Operation = statemanager.BASE

	h.api.Send(msg)

}
