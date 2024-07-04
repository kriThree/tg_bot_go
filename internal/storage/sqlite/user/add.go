package user

import (
	"context"
	"english_learn/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

func (s User) Add(ctx context.Context, tgId int) (string, error) {

	const op = "storage.user.Add"

	stmt, err := s.db.Prepare("INSERT INTO users (id, tg_id) VALUES (?, ?)")

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := uuid.New().String()

	_, err = stmt.ExecContext(
		ctx,
		id,
		tgId,
	)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, storage.InternalError)
	}

	return id, nil
}
