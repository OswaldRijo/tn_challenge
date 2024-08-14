package operations

import (
	"context"
	"fmt"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
)

func (u *OperationsApiImpl) GetUserBalance(ctx context.Context, userId int64) (*operationspb.Balance, error) {
	userBalance, err := u.balancesRepo.GetBalanceByUserId(ctx, userId)
	if err != nil {
		return nil, common.NewAPIErrorInternal(err)
	}

	if userBalance == nil {
		return nil, common.NewAPIErrorResourceNotFound(fmt.Errorf(UserBalanceNotFound))
	}

	return ParseBalanceModelToPb(userBalance), nil
}
