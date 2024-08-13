package balances

import (
	"context"

	"gorm.io/gorm"

	"truenorth/packages/database"
	"truenorth/services/operations_service/models"
)

//go:generate mockery --name=BalancesRepo --output=../../../../mocks/operations_service/repositories
type BalancesRepo interface {
	GetBalanceByUserId(ctx context.Context, userId int64) (*models.Balance, error)
	CreateBalance(ctx context.Context, balance *models.Balance, tx *gorm.DB) error
	UpdateBalance(ctx context.Context, balance *models.Balance, tx *gorm.DB) error
}

type BalancesRepoImpl struct {
	db *gorm.DB
}

func NewBalancesRepo() BalancesRepo {
	return &BalancesRepoImpl{
		db: database.GetInstance(),
	}
}
