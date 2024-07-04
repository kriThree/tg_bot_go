package botapp

import (
	"english_learn/internal/bot"
	"english_learn/internal/bot/controller"
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

func New(log *slog.Logger, token string, controller * controller.Controller) *App {
	botServer := bot.New(token, log, controller)

	return &App{token: token, l: log, botServer: botServer}
}
func (app *App) Run() error {
	app.botServer.Run()
	return nil
}
