package operations

import (
	"context"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (ori *OperationsRepoImpl) CreateOperation(ctx context.Context, operation *models.Operation, tx *gorm.DB) error {
	result := tx.WithContext(ctx).Create(operation)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
