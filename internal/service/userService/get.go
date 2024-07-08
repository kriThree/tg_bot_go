package user_service

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"
)

func (s UserService) GetUserByTgId(ctx context.Context, tgId int) (models.User, error) {

	const op = "user_service.get"

	log := s.l.With(slog.String("op", op))

	user, err := s.up.GetByTgId(ctx, tgId)

	if err != nil {
		log.Error("Get user error", slog.Any("error", err))
		return models.User{}, nil
	}

	return user, nil
}
