package definition_service

import (
	"context"
	"english_learn/internal/domain/models"
	"fmt"
)

func (s *BotService) AddDefinition(ctx context.Context, definitionWord string, meaning models.Meaning) (string, error) {

	const op = "definition_sevice.add"

	defId, err := s.dp.Add(ctx, definitionWord)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, err = s.mp.Add(ctx, meaning, defId)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return defId, nil
}
