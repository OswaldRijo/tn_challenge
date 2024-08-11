package database

import (
	"context"
	"database/sql"
)

func RunDbTransaction(ctx context.Context, onRunningTxn func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // nolint:errcheck
	err = onRunningTxn(ctx, tx)
	if err != nil {
		return err
	}
	return tx.Commit()
}
