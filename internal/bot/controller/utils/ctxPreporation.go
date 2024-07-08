package utils

import (
	"context"
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CtxPreporation(ctx context.Context, update tgbotapi.Update, state *statemanager.State) (int, AppContext, error) {

	var id int
	if update.Message != nil {
		id = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		id = update.CallbackQuery.From.ID
	}

	if id == 0 {
		return 0, AppContext{}, NotValidUpdateErr
	}
	if state.GetUser(id) == (statemanager.UserState{}) {
		state.SetUser(id, statemanager.UserState{
			Operation: statemanager.BASE,
			TgID:      id,
		})
	}

	uState := state.GetUser(id)

	return id, AppContext{Update: update, State: &uState}, nil

}
