package operations

import (
	"context"

	"truenorth/packages/common"
	operationspb "truenorth/pb/operations"
)

func (u *OperationsApiImpl) FilterRecords(ctx context.Context, operationReq *operationspb.FilterRecordsRequest) ([]*operationspb.Record, int64, error) {
	records, totalCount, err := u.recordsRepo.FindRecordsByUserId(ctx, operationReq.GetUserId(), operationReq.GetLimit(), operationReq.GetPage(), operationReq.GetOrderByFields()...)

	if err != nil {
		return nil, 0, common.NewAPIErrorInternal(err)
	}

	return ParseRecordModelArrToPb(records), totalCount, nil
}
