package definition_sevice

import (
	"context"
	"english_learn/internal/domain/models"
)

func (s *BotService) AddDefinition(ctx context.Context, definition models.Definition) (string, error) {

	return s.dp.AddDefinition(ctx, definition)
}
