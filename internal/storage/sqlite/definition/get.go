package definition

import (
	"context"
	"database/sql"
	"english_learn/internal/domain/models"
	"english_learn/internal/storage"
	"errors"
	"fmt"
)

func (d Definition) GetDefinition(ctx context.Context, id string) (models.Definition, error) {

	const op = "storage.definition.Get"

	stmt, err := d.db.Prepare("SELECT id, word, created_at FROM definitions WHERE id = ?")

	if err != nil {
		return models.Definition{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, id)

	var definition models.Definition

	err = row.Scan(&definition.ID, &definition.Word, &definition.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Definition{}, fmt.Errorf("%s: %w", op, storage.DefinitionNotFound)
		}
		return models.Definition{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.Definition{}, nil
}
