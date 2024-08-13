package records

import (
	"context"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) CreateRecord(ctx context.Context, record *models.Record, tx *gorm.DB) error {
	result := tx.WithContext(ctx).Create(record)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
