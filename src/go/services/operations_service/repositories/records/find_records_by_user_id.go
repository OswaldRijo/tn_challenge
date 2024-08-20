package records

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/pb/operations"
	"truenorth/services/operations_service/models"
)

func (rri *RecordsRepoImpl) FindRecordsByUserId(ctx context.Context, userId int64, limit, page int32, orderBy ...*operations.OrderBy) ([]*models.Record, int64, error) {
	var records []*models.Record
	var totalCount int64
	q := rri.db.WithContext(ctx).Preload("Operation").Where(map[string]interface{}{"user_id": userId, "deleted": false}).
		Offset(int(limit * page)).Limit(int(limit))
	q = setUpOrderBy(q, orderBy...)
	result := q.Find(&records)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, result.Error
	}

	result = rri.db.WithContext(ctx).Model(&models.Record{}).Where(map[string]interface{}{"user_id": userId, "deleted": false}).Count(&totalCount)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, result.Error
	}

	return records, totalCount, nil

}

func setUpOrderBy(q *gorm.DB, orderBy ...*operations.OrderBy) *gorm.DB {

	for _, o := range orderBy {
		switch o.OrderField {
		case operations.OrderFieldsEnum_ID:
			if o.OrderType == operations.OrderTypeEnum_ASC {
				q.Order("id asc")
			} else {
				q.Order("id desc")
			}
			break
		case operations.OrderFieldsEnum_CREATED_AT:
			if o.OrderType == operations.OrderTypeEnum_ASC {
				q.Order("created_at asc")
			} else {
				q.Order("created_at desc")
			}
			break
		case operations.OrderFieldsEnum_OPERATION:
			if o.OrderType == operations.OrderTypeEnum_ASC {
				q.Order("operation_type asc")
			} else {
				q.Order("operation_type desc")
			}
			break
		}
	}
	return q
}
