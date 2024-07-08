package user_service

import (
	"context"
	"english_learn/internal/domain/models"
	"fmt"
	"log/slog"
)

func (s UserService) UserTgInteraction(ctx context.Context, tgId int) (models.User, error) {

	const op = "user_service.tg_interaction"

	log := s.l.With(slog.String("op", op))

	user, err := s.up.GetByTgId(ctx, tgId)

	if err != nil {
		log.Error("Get user error", slog.Any("error", err))

		id, err := s.up.Add(ctx, tgId)

		if err != nil {
			log.Error("Add user error", slog.Any("error", err))

			return models.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return models.User{ID: id, TGID: tgId}, nil
	}

	return user, nil
}
