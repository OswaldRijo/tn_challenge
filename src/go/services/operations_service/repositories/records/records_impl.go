package records

import (
	"context"

	"gorm.io/gorm"

	"truenorth/packages/database"
	"truenorth/services/operations_service/models"
)

//go:generate mockery --name=RecordsRepo --output=../../../../mocks/operations_service/repositories
type RecordsRepo interface {
	FindRecordsByUserId(ctx context.Context, userId int64, limit, page int) ([]*models.Record, error)
	CreateRecord(ctx context.Context, operation *models.Record, tx *gorm.DB) error
	DeleteRecordById(ctx context.Context, recordId int64, tx *gorm.DB) error
	FindRecordsByIds(ctx context.Context, ids ...int64) ([]*models.Record, error)
}

type RecordsRepoImpl struct {
	db *gorm.DB
}

func NewRecordsRepo() RecordsRepo {
	return &RecordsRepoImpl{
		db: database.GetInstance(),
	}
}
