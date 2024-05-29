package definition

import "database/sql"

type Definition struct {
	db *sql.DB
}

func New(db *sql.DB) Definition {
	return Definition{db: db}
}
