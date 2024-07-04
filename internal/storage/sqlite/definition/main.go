package definition

import (
	sqlite_utils "english_learn/internal/storage/sqlite/utils"
)

type Definition struct {
	db *sqlite_utils.AppDB
}

func New(db *sqlite_utils.AppDB) Definition {
	return Definition{db: db}
}
