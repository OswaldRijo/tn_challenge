package records

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) FindRecordsByUserId(ctx context.Context, userId int64, limit, page int) ([]*models.Record, error) {
	records := []*models.Record{}
	result := rri.db.WithContext(ctx).Where(map[string]interface{}{"user_id": userId}).Find(&records).
		Offset(limit * page).Limit(limit)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return records, nil

}
