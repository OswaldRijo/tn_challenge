package records

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) FindRecordsByIds(ctx context.Context, ids ...int64) ([]*models.Record, error) {
	records := []*models.Record{}
	result := rri.db.WithContext(ctx).Where(map[string]interface{}{"id": ids}).Find(&records)
	
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return records, nil

}
