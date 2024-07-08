package user

import (
	"context"
	"database/sql"
	"english_learn/internal/domain/models"
	storage "english_learn/internal/storage"
	"errors"
	"fmt"
)

func (s User) GetById(ctx context.Context, id string) (models.User, error) {
	const op = "storage.user.Get"

	stmt, err := s.db.Prepare("SELECT id, tg_id FROM users WHERE id = ?")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, id)

	var user models.User
	err = row.Scan(&user.ID, &user.TGID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.UserNotFoundErr)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}

func (s User) GetByTgId(ctx context.Context, tgId int) (models.User, error) {
	const op = "sqlite.user.GetByTgId"

	stmt, err := s.db.Prepare("SELECT id, tg_id FROM users WHERE tg_id = ?")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, tgId)

	var user models.User
	err = row.Scan(&user.ID, &user.TGID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.UserNotFoundErr)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}
