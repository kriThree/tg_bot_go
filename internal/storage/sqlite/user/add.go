package user

import (
	"context"
	"english_learn/internal/storage"
	"errors"
	"fmt"

	sqlite3 "github.com/mattn/go-sqlite3"

	"github.com/google/uuid"
)

func (s User) Add(ctx context.Context, tgId int) (string, error) {

	const op = "sqlite.user.Add"

	stmt, err := s.db.Prepare("INSERT INTO users (id, tg_id) VALUES (?, ?)")

	if err != nil {
		if errors.Is(err, sqlite3.ErrConstraintUnique) {
			return "", fmt.Errorf("%s: %w", op, storage.UserAlreadyAddedErr)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := uuid.New().String()

	_, err = stmt.ExecContext(
		ctx,
		id,
		tgId,
	)

	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code == sqlite3.ErrNo(sqlite3.ErrConstraint) {
			return "", fmt.Errorf("%s: %w", op, storage.UserAlreadyAddedErr)
		}
		return "", fmt.Errorf("%s: %w, %w", op, err, storage.InternalErr)
	}

	return id, nil
}
