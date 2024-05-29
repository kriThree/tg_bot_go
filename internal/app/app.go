package app

import (
	botapp "english_learn/internal/app/bot"
	definition_sevice "english_learn/internal/service/definition"
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

	botService := definition_sevice.New(log, storage.Definition)

	botApp := botapp.New(log, token, botService)

	return &App{
		bot: *botApp,
	}
}

func (app *App) MustRun() {
	if err := app.bot.Run(); err != nil {
		panic(err)
	}
}
