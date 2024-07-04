package controller

import (
	"context"
	handlers "english_learn/internal/bot/controller/handlers"
	"english_learn/internal/bot/controller/middlewares"
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Router func(ctx context.Context, update tgbotapi.Update, uState *statemanager.UserState)

func createRouter(h *handlers.BotHandlers, m *middlewares.Middleware) Router {
	return func(ctx context.Context, update tgbotapi.Update, uState *statemanager.UserState) {

		ctx = context.Background()

		if uState.Operation == "" {
			h.Start(update, uState)
			return
		}
		if update.Message == nil {
			if update.CallbackQuery != nil {
				switch update.CallbackQuery.Data {
				case handlers.SAVE_DEFINITION_END:
					h.SaveDefinitionEnd(ctx, update, uState)
				}
			}
			return
		} else {
			switch update.Message.Text {
			case handlers.SAVE_DEFINITION_QUERY:
				h.SaveDefinitionQuery(ctx, update, uState)
			default:
				switch uState.Operation {
				case statemanager.SAVE_DEFINITION_WAIT_NAME:
					h.SaveDefinitionAddName(ctx, update, uState)
				case statemanager.SAVE_DEFINITION_WAIT_MEANING:
					h.SaveDefinitionAddMean(ctx, update, uState)
				default:
					h.Start(update, uState)
				}
			}
		}
	}
}
