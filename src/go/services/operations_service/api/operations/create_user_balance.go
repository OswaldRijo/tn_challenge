package operations

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"truenorth/packages/database"
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

	err = database.PerformDbTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		now := time.Now()
		userBalance = models.NewBalance().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetUserID(userId).
			SetCurrentBalance(config.Config.DefaultUserBalance)
		err = u.balancesRepo.CreateBalance(ctx, userBalance, tx)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		return nil
	})

	return ParseBalanceModelToPb(userBalance), nil
}
