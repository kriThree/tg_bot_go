package user_service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s UserService) AddUser(ctx context.Context, tgId int) (string, error) {

	const op = "user_service.add"

	log := s.l.With(slog.String("op", op))

	id, err := s.up.Add(ctx, tgId)

	if err != nil {
		log.Error("Add user error", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil

}
