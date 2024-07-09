package definition

import (
	"context"
	storage "english_learn/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

func (s Definition) Add(ctx context.Context, word string, userId string) (string, error) {

	const op = "sqlite.definition.Add"

	stmt, err := s.db.Prepare("INSERT INTO definitions (id, word, user_id) VALUES (?, ?, ?)")

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := uuid.New().String()

	_, err = stmt.ExecContext(
		ctx,
		id,
		word,
		userId,
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, storage.InternalErr)
	}

	return id, nil
}
