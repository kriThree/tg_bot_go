package meaning

import (
	"context"
	"english_learn/internal/domain/models"
	"english_learn/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

func (s Meaning) Add(
	ctx context.Context,
	meaning models.Meaning,
	definitionId string,
) (string, error) {

	const op = "storage.meaning.Add"

	stmt, err := s.db.Prepare("INSERT INTO meanings (id, definition_id, part_of_speach,value) VALUES (?, ?, ?, ?)")

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id := uuid.New().String()

	_, err = stmt.ExecContext(
		ctx,
		id,
		definitionId,
		meaning.PartOfSpeach,
		meaning.Value,
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, storage.InternalError)
	}

	return id, nil
}
