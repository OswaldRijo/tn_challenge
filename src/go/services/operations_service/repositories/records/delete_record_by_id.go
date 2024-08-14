package records

import (
	"context"
	"time"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) DeleteRecordById(ctx context.Context, recordId int64, tx *gorm.DB) error {
	records := &models.Record{}
	result := tx.WithContext(ctx).Preload("Operation").Where(map[string]interface{}{"id": recordId, "deleted": false}).First(&records)
	if result.Error != nil {
		return result.Error
	}

	records.Deleted = true
	records.UpdatedAt = time.Now()
	return tx.Save(records).Error
}
