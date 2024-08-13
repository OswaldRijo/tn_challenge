package balances

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (bri *BalancesRepoImpl) GetBalanceByUserId(ctx context.Context, userId int64) (*models.Balance, error) {
	balance := &models.Balance{}
	result := bri.db.WithContext(ctx).Where(map[string]interface{}{"user_id": userId}).First(balance)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return balance, nil
}
