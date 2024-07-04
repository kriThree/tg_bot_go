package utils

import (
	"context"
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func UserInApp(ctx context.Context, update tgbotapi.Update, state *statemanager.State) (int, error) {

	var id int
	if update.Message != nil {
		id = update.Message.From.ID
	} else if update.CallbackQuery != nil {
		id = update.CallbackQuery.From.ID
	}

	if id == 0 {
		return 0, NotValidUpdateErr
	}
	if state.GetUser(id) == (statemanager.UserState{}) {
		state.SetUser(id, statemanager.UserState{
			Operation: statemanager.BASE,
			ID:        id,
		})
	}

	return id, nil

}
