package handlers

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DefinitionUscase interface {
	AddDefinition(ctx context.Context, definition models.Definition) (string, error)
	GetDefinition(ctx context.Context, id string) (models.Definition, error)
}

type BotHandlers struct {
	l   *slog.Logger
	api *tgbotapi.BotAPI
	du  DefinitionUscase
}

func New(l *slog.Logger, du DefinitionUscase, api *tgbotapi.BotAPI) *BotHandlers {
	return &BotHandlers{
		l:   l,
		du:  du,
		api: api,
	}
}
