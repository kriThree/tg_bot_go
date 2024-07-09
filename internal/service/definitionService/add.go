package definition_service

import (
	"context"
	"english_learn/internal/domain/models"
	"fmt"
	"log/slog"
)

func (s *BotService) AddDefinition(
	ctx context.Context,
	definitionWord string,
	meaning models.Meaning,
	userId string,
) (string, error) {

	const op = "definition_service.add"

	log := s.l.With(slog.String("op", op))

	defId, err := s.dp.Add(ctx, definitionWord, userId)

	if err != nil {
		log.Error("Add definition error", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, err = s.mp.Add(ctx, meaning, defId)

	if err != nil {
		log.Error("Add meaning error", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return defId, nil
}
