package botapp

import (
	"english_learn/internal/bot"
	botHandlers "english_learn/internal/bot/handlers"
	"log/slog"
)

type BotServer interface {
	Run() error
}

type App struct {
	l         *slog.Logger
	token     string
	botServer BotServer
}

func New(log *slog.Logger, token string, botUsecase botHandlers.DefinitionUscase) *App {
	botServer := bot.New(token, log, botUsecase)

	return &App{token: token, l: log, botServer: botServer}
}
func (app *App) Run() error {
	app.botServer.Run()
	return nil
}
