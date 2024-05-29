package definition

import (
	"context"
	"english_learn/internal/domain/models"
	storage "english_learn/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

func (s Definition) AddDefinition(ctx context.Context, definition models.Definition) (string, error) {

	const op = "storage.definition.Add"

	stmt, err := s.db.Prepare("INSERT INTO definitions (id, word, created_at) VALUES (?, ?, ?)")

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := uuid.New().String()

	_, err = stmt.ExecContext(
		ctx,
		id,
		definition.Word,
		definition.CreatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, storage.InternalError)
	}

	return id, nil
}
