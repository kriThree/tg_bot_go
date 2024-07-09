package controller

import (
	handlers "english_learn/internal/bot/controller/handlers"
	"english_learn/internal/bot/controller/middlewares"
	"english_learn/internal/bot/controller/utils"
	statemanager "english_learn/internal/bot/stateManager"
	"fmt"
	"strings"
)

func createRouter(h *handlers.BotHandlers, m *middlewares.BotMiddlewares) utils.Handler {

	middlewarer := createMiddlewarer(m.PutUser)

	return func(ctx utils.AppContext) {

		err := middlewarer(ctx)

		if err != nil {
			return
		}

		if ctx.State.Operation == "" {
			h.Start(ctx)
			return
		}

		if ctx.Update.Message == nil {
			fmt.Println("message is nil", ctx.Update.CallbackQuery, ctx.Update.CallbackQuery.Data, strings.Contains(ctx.Update.CallbackQuery.Data, handlers.SAVE_DEFINITION_END))
			if ctx.Update.CallbackQuery != nil {
				if strings.Contains(ctx.Update.CallbackQuery.Data, handlers.SAVE_DEFINITION_END) {
					h.SaveDefinitionEnd(ctx)
				}
			}
			return
		} else {
			switch ctx.Update.Message.Text {
			case handlers.SAVE_DEFINITION_QUERY:
				h.SaveDefinitionQuery(ctx)
			default:
				switch ctx.State.Operation {
				case statemanager.SAVE_DEFINITION_WAIT_NAME:
					h.SaveDefinitionAddName(ctx)
				case statemanager.SAVE_DEFINITION_WAIT_MEANING:
					h.SaveDefinitionAddMean(ctx)
				default:
					h.Start(ctx)
				}
			}
		}
	}
}

type Middleware func(ctx utils.AppContext) error

func createMiddlewarer(middlewares ...Middleware) Middleware {
	return func(ctx utils.AppContext) error {
		for _, middleware := range middlewares {

			if err := middleware(ctx); err != nil {
				return err
			}

			return nil
		}
		return nil
	}

}
