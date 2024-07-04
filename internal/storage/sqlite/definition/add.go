package definition

import (
	"context"
	storage "english_learn/internal/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s Definition) Add(ctx context.Context, word string) (string, error) {

	const op = "storage.definition.Add"

	stmt, err := s.db.Prepare("INSERT INTO definitions (id, word, created_at) VALUES (?, ?, ?)")

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := uuid.New().String()

	_, err = stmt.ExecContext(
		ctx,
		id,
		word,
		time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, storage.InternalError)
	}

	return id, nil
}
