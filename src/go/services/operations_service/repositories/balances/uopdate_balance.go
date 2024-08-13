package balances

import (
	"context"

	"gorm.io/gorm"

	"truenorth/services/operations_service/models"
)

func (bri *BalancesRepoImpl) UpdateBalance(ctx context.Context, balance *models.Balance, tx *gorm.DB) error {
	result := tx.WithContext(ctx).Updates(balance)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
