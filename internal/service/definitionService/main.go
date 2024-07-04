package definition_service

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"
)

type BotService struct {
	l  *slog.Logger
	dp DefinitionProvider
	mp MeaningProvider
}
type DefinitionProvider interface {
	Add(ctx context.Context, word string) (string, error)
	Get(ctx context.Context, id string) (models.Definition, error)
	GetMany(ctx context.Context) ([]models.Definition, error)
}
type MeaningProvider interface {
	Add(ctx context.Context, meaning models.Meaning, definitionId string) (string, error)
	// Get(ctx context.Context, id string) (models.Meaning, error)
	// Delete(ctx context.Context, id string) error
}
//Определение интерфйса по месту объявления для разгрузки кода


func New(l *slog.Logger, dp DefinitionProvider, mp MeaningProvider) *BotService {
	return &BotService{l: l, dp: dp, mp: mp}
}
