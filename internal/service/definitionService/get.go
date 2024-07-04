package definition_service

import (
	"context"
	"english_learn/internal/domain/models"
)

func (s *BotService) GetDefinition(ctx context.Context, word string) (models.Definition, error) {

	definition, err := s.dp.Get(ctx, word)

	if err != nil {
		return models.Definition{}, err
	}

	return definition, nil
}
