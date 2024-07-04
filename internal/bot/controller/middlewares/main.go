package middlewares

import (
	"context"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotMiddlewares struct {
	l   *slog.Logger
	api *tgbotapi.BotAPI
	mu  MiddlewaresUsecase
}

type MiddlewaresUsecase interface {
	IsUserExist(ctx context.Context, id int) (bool, error)
	AddUser(ctx context.Context, id int) (string, error)
}

func New() *BotMiddlewares {
	return &BotMiddlewares{}
}
