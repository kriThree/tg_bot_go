package bot

import (
	handlers "english_learn/internal/bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func createRouter(h *handlers.BotHandlers) func(update tgbotapi.Update) {
	return func(update tgbotapi.Update) {
		if update.Message != nil {
			h.SaveDefinition(update)
		}
	}
}
