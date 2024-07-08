package controller

import (
	"context"
	"english_learn/internal/bot/controller/handlers"
	"english_learn/internal/bot/controller/middlewares"
	"english_learn/internal/bot/controller/utils"
	statemanager "english_learn/internal/bot/stateManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"log/slog"
)

type Controller struct {
	state *statemanager.State
	hu    handlers.HandlersUsecase
	mu    middlewares.MiddlewaresUsecase
	l     *slog.Logger
}

func New(l *slog.Logger, hu handlers.HandlersUsecase, mu middlewares.MiddlewaresUsecase) *Controller {
	return &Controller{
		state: statemanager.New(),
		hu:    hu,
		l:     l,
		mu:    mu,
	}

}
func (c *Controller) Handle(updates <-chan tgbotapi.Update, api *tgbotapi.BotAPI) {

	log := c.l.With(slog.String("op", "controller.Handle"))

	state := statemanager.New()

	router := createRouter(handlers.New(c.l, c.hu, api), middlewares.New(c.l, c.mu, api))

	for update := range updates {

		ctx := context.Background()

		id, appCtx, err := utils.CtxPreporation(ctx, update, state)

		if err != nil {
			log.Error("User error", slog.Any("error", err))
			continue
		}

		if update.Message != nil {
			router(appCtx)
			state.SetUser(id, *appCtx.State)
		}
		log.Info("State", slog.Any("state", state))

	}
}
