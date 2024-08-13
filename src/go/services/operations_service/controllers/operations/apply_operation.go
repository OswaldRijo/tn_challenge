package operations

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
)

func (uc *OperationsControllerImpl) ApplyOperation(ctx context.Context, req *operationspb.ApplyOperationRequest) (*operationspb.ApplyOperationResponse, error) {
	operation, record, balance, err := uc.operationsApi.ApplyOperation(ctx, req)

	if err != nil {
		return nil, status.Error(common.HandleApiError(err), err.Error())
	}
	return &operationspb.ApplyOperationResponse{
		Operation:          operation,
		CurrentUserBalance: balance,
		Record:             record,
	}, nil
}
