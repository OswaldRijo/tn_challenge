package records

import (
	"context"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) DeleteRecordById(ctx context.Context, recordId int64, tx *gorm.DB) error {
	operation := &models.Record{ID: recordId, Deleted: true}
	result := tx.WithContext(ctx).Save(operation)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
