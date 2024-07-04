package sqlite

import (
	"english_learn/internal/storage/sqlite/definition"
	"english_learn/internal/storage/sqlite/meaning"
	"english_learn/internal/storage/sqlite/user"
	sqlite_utils "english_learn/internal/storage/sqlite/utils"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db         *sqlite_utils.AppDB
	Definition definition.Definition
	Meaning    meaning.Meaning
	User       user.User
}

func New(storagePath string) (*Storage, error) {

	const op = "storage.sqlite.New"

	db, err := sqlite_utils.Open("sqlite3", storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{
		db:         db,
		Definition: definition.New(db),
		Meaning:    meaning.New(db),
		User:       user.New(db),
	}, nil
}
