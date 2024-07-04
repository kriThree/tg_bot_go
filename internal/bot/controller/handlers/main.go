package handlers

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotHandlers struct {
	api *tgbotapi.BotAPI
	hu  HandlersUsecase
	l   *slog.Logger
}

type HandlersUsecase interface {
	AddDefinition(ctx context.Context, definitionWord string, meaning models.Meaning) (string, error)
	GetDefinition(ctx context.Context, id string) (models.Definition, error)
	GetDefinitions(ctx context.Context) ([]models.Definition, error)
}

func New(l *slog.Logger, hu HandlersUsecase, api *tgbotapi.BotAPI) *BotHandlers {
	return &BotHandlers{
		l:   l,
		hu:  hu,
		api: api,
	}
}
