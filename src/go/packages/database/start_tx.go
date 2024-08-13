package database

import (
	"context"

	"gorm.io/gorm"
)

func PerformDbTransaction(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	err := onRunningTxn(ctx, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
