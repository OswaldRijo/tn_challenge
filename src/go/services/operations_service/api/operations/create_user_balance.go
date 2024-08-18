package operations

import (
	"context"
	"time"

	"gorm.io/gorm"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/config"
	"truenorth/services/operations_service/models"
)

func (u *OperationsApiImpl) CreateUserBalance(ctx context.Context, userId int64) (*operationspb.Balance, error) {
	userBalance, err := u.balancesRepo.GetBalanceByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if userBalance != nil {
		return ParseBalanceModelToPb(userBalance), nil
	}

	err = InitTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		now := time.Now()
		userBalance = models.NewBalance().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetUserID(userId).
			SetCurrentBalance(config.Config.DefaultUserBalance)
		err = u.balancesRepo.CreateBalance(ctx, userBalance, tx)
		if err != nil {
			return common.NewAPIErrorInternal(err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ParseBalanceModelToPb(userBalance), nil
}
