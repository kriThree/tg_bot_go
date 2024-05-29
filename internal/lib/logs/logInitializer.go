package log

import (
	baseLog "log"
	"log/slog"
	"os"
)

const (
	env_local = "local"
	env_prod  = "prod"
	env_dev   = "dev"
)

// Логгеры
// Для дебага лучше использовать log или fmt

var Log baseLog.Logger
var FileLog baseLog.Logger
var ConsoleLog baseLog.Logger

func LogInitializer(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case env_local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case env_dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case env_prod:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log

}
