package operations

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
)

func (uc *OperationsControllerImpl) DeleteRecords(ctx context.Context, req *operationspb.DeleteRecordsRequest) (*operationspb.DeleteRecordsResponse, error) {
	record, balance, err := uc.operationsApi.DeleteRecord(ctx, req)

	if err != nil {
		return nil, status.Error(common.HandleApiError(err), err.Error())
	}
	return &operationspb.DeleteRecordsResponse{
		Records:        record,
		CurrentBalance: balance,
	}, nil
}
