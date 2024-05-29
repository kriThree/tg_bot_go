package definition_sevice

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"
)

type BotService struct {
	l  *slog.Logger
	dp DefinitionProvider
}
type DefinitionProvider interface {
	AddDefinition(ctx context.Context, definition models.Definition) (string, error)
	GetDefinition(ctx context.Context, id string) (models.Definition, error)
	GetManyDefinitionsByDate(ctx context.Context, date string) ([]models.Definition, error)
}
type MeaningProvider interface {
	// Add(ctx context.Context, meaning models.Meaning, definitionId string) (string, error)
	// Get(ctx context.Context, id string) (models.Meaning, error)
	// Delete(ctx context.Context, id string) error
}

func New(l *slog.Logger, dp DefinitionProvider) *BotService {
	return &BotService{l: l, dp: dp}
}
