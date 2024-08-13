package operations

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (ori *OperationsRepoImpl) FindRecordsByUserId(ctx context.Context, userId int64, limit, page int) ([]*models.Operation, error) {
	operations := []*models.Operation{}
	result := ori.db.WithContext(ctx).Where(map[string]interface{}{"user_id": userId}).Find(&operations).
		Offset(limit * page).Limit(limit)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return operations, nil

}
