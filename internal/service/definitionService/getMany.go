package definition_service

import (
	"context"
	"english_learn/internal/domain/models"
)

func (s *BotService) GetDefinitions(ctx context.Context) ([]models.Definition, error) {
	definitions, err := s.dp.GetMany(ctx)

	if err != nil {
		return []models.Definition{}, err
	}

	return definitions, nil

}
