package operations

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
)

func (uc *OperationsControllerImpl) GetUserBalance(ctx context.Context, req *operationspb.GetUserBalanceRequest) (*operationspb.GetUserBalanceResponse, error) {
	balance, err := uc.operationsApi.GetUserBalance(ctx, req.GetUserId())

	if err != nil {
		return nil, status.Error(common.HandleApiError(err), err.Error())
	}
	return &operationspb.GetUserBalanceResponse{
		Balance: balance,
	}, nil
}
