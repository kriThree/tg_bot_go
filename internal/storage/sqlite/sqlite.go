package sqlite

import (
	"database/sql"
	"english_learn/internal/storage/sqlite/definition"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db         *sql.DB
	Definition definition.Definition
}

func New(storagePath string) (*Storage, error) {

	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{
		db:         db,
		Definition: definition.New(db),
	}, nil
}
