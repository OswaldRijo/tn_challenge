package operations

import (
	"context"

	"gorm.io/gorm"

	"truenorth/packages/database"
	"truenorth/services/operations_service/models"
)

//go:generate mockery --name=OperationsRepo --output=../../../../mocks/operations_service/repositories
type OperationsRepo interface {
	CreateOperation(ctx context.Context, operation *models.Operation, tx *gorm.DB) error
}

type OperationsRepoImpl struct {
	db *gorm.DB
}

func NewOperationsRepo() OperationsRepo {
	return &OperationsRepoImpl{
		db: database.GetInstance(),
	}
}
