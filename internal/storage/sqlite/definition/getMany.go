package definition

import (
	"context"
	"english_learn/internal/domain/models"
	"fmt"
)

func (s Definition) getManyByFilter(ctx context.Context, filter string, args ...interface{}) ([]models.Definition, error) {

	const op = "sqlite.definition.GetMany"

	stmt, err := s.db.Prepare("SELECT id, word, created_at FROM definitions " + filter)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var definitions []models.Definition

	rows, err := stmt.QueryContext(ctx, args...)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		var definition models.Definition

		err = rows.Scan(&definition.ID, &definition.Word, &definition.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		definitions = append(definitions, definition)
	}

	return definitions, nil
}

func (s Definition) GetMany(ctx context.Context) ([]models.Definition, error) {

	return s.getManyByFilter(ctx, "WHERE created_at = ?")

}
