package records

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) FindRecordsByUserId(ctx context.Context, userId int64, limit, page int32) ([]*models.Record, int64, error) {
	var records []*models.Record
	var totalCount int64
	result := rri.db.WithContext(ctx).Preload("Operation").Where(map[string]interface{}{"user_id": userId, "deleted": false}).
		Offset(int(limit * page)).Limit(int(limit)).Find(&records)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, result.Error
	}

	result = rri.db.WithContext(ctx).Model(&models.Record{}).Where(map[string]interface{}{"user_id": userId, "deleted": false}).Count(&totalCount)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, result.Error
	}

	return records, totalCount, nil

}
