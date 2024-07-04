package meaning

import (
	sqlite_utils "english_learn/internal/storage/sqlite/utils"
)

type Meaning struct {
	db *sqlite_utils.AppDB
}

func New(db *sqlite_utils.AppDB) Meaning {
	return Meaning{db: db}
}
