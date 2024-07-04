package user

import (
	sqlite_utils "english_learn/internal/storage/sqlite/utils"
)

type User struct {
	db *sqlite_utils.AppDB
}

func New(db *sqlite_utils.AppDB) User {
	return User{db: db}
}
