package user_service

import (
	"context"
	"english_learn/internal/domain/models"
	"log/slog"
)

type UserService struct {
	l  *slog.Logger
	up UserProvider
}

type UserProvider interface {
	Add(ctx context.Context, tgId int) (string, error)
	GetByTgId(ctx context.Context, id int) (models.User, error)
}

func New(l *slog.Logger, up UserProvider) *UserService {
	return &UserService{l: l, up: up}
}
