package sqlite_utils

import (
	"context"
	"database/sql"
	"fmt"
)

type txKey struct{}

func InjectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}

func (db *AppDB) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Commit()

	err = tFunc(InjectTx(ctx, tx))
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func Open(driverName, dataSourceName string) (*AppDB, error) {

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return &AppDB{}, err
	}

	return &AppDB{db}, db.Ping()

}

type AppDB struct {
	*sql.DB
}
