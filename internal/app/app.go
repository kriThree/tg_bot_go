package app

import (
	botapp "english_learn/internal/app/bot"
	"english_learn/internal/bot/controller"
	definition_service "english_learn/internal/service/definitionService"
	user_service "english_learn/internal/service/userService"
	"english_learn/internal/storage/sqlite"
	"log/slog"
)

type App struct {
	bot botapp.App
}

func New(log *slog.Logger, token string, storagePath string) *App {

	storage, err := sqlite.New(storagePath)

	if err != nil {
		panic(err)
	}

	definitionService := definition_service.New(log, storage.Definition, storage.Meaning)
	
	userService := user_service.New(log, storage.User)

	controller := controller.New(log, definitionService, userService)

	botApp := botapp.New(log, token, controller)

	return &App{
		bot: *botApp,
	}
}

func (app *App) MustRun() {
	if err := app.bot.Run(); err != nil {
		panic(err)
	}
}
