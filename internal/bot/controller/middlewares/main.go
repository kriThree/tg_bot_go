package middlewares

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotMiddlewares struct {
	l   *slog.Logger
	api *tgbotapi.BotAPI
	mu  MiddlewaresUsecase
}

type MiddlewaresUsecase interface {
	UserTgInteraction(ctx context.Context, id int) (models.User, error)
}

func New(l *slog.Logger, mu MiddlewaresUsecase, api *tgbotapi.BotAPI) *BotMiddlewares {
	return &BotMiddlewares{l: l, mu: mu, api: api}
}
