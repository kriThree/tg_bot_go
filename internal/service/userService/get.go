package user_service

import (
	"context"
	"log/slog"
)

func (s UserService) IsUserExist(ctx context.Context, id int) (bool, error) {

	const op = "user_service.get"

	log := s.l.With(slog.String("op", op))

	_, err := s.up.GetByTgId(ctx, id)

	if err != nil {
		log.Error("Get user error", slog.Any("error", err))
		return false, nil
	}

	return true, nil

}
