package operations_test

import (
	"errors"

	"google.golang.org/grpc/status"

	operationspb "truenorth/pb/operations"
	"truenorth/services/users_service/test"
)

func (ucts *OperationControllerTestSuite) Test_FilterRecords_Success() {
	// arrange
	count := int64(1)
	userId := int64(1)
	limit := int32(1)
	page := int32(1)
	req := &operationspb.FilterRecordsRequest{UserId: &userId, Limit: &limit, Page: &page}
	ucts.operationsApiMock.On("FilterRecords", ucts.ctx, req).Return(test.RecordsPb, count, nil).Once()

	// act
	res, err := ucts.controller.FilterRecords(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.RecordsPb, res.GetRecords())
	ucts.Equal(count, res.GetTotalCount())
	ucts.assertAllMocks()
}

func (ucts *OperationControllerTestSuite) Test_FilterRecords_ApiError() {
	// arrange
	userId := int64(1)
	limit := int32(1)
	page := int32(1)
	errExpected := errors.New("error getting user balance")
	req := &operationspb.FilterRecordsRequest{UserId: &userId, Limit: &limit, Page: &page}
	ucts.operationsApiMock.On("FilterRecords", ucts.ctx, req).Return(nil, int64(0), errExpected).Once()

	// act
	res, err := ucts.controller.FilterRecords(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected.Error(), rpcStatus.Message())
	ucts.assertAllMocks()
}
