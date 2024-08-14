package operations

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
)

func (uc *OperationsControllerImpl) FilterRecords(ctx context.Context, req *operationspb.FilterRecordsRequest) (*operationspb.FilterRecordsResponse, error) {
	records, totalCount, err := uc.operationsApi.FilterRecords(ctx, req)

	if err != nil {
		return nil, status.Error(common.HandleApiError(err), err.Error())
	}
	return &operationspb.FilterRecordsResponse{
		Records:    records,
		TotalCount: totalCount,
	}, nil
}
