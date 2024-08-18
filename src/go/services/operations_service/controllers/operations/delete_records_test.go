package operations_test

import (
	"errors"

	"google.golang.org/grpc/status"

	operationspb "truenorth/pb/operations"
	"truenorth/services/users_service/test"
)

func (ucts *OperationControllerTestSuite) Test_DeleteRecords_Success() {
	// arrange
	recordId := int64(1)
	userId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: userId}
	ucts.operationsApiMock.On("DeleteRecord", ucts.ctx, req).Return(test.RecordsPb, test.YagamiLightBalancePb, nil).Once()

	// act
	res, err := ucts.controller.DeleteRecords(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.RecordsPb, res.GetRecords())
	ucts.Equal(test.YagamiLightBalancePb, res.GetCurrentBalance())
	ucts.assertAllMocks()
}

func (ucts *OperationControllerTestSuite) Test_DeleteRecords_ApiError() {
	// arrange
	recordId := int64(1)
	userId := int64(1)
	errExpected := errors.New("error getting user balance")
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: userId}
	ucts.operationsApiMock.On("DeleteRecord", ucts.ctx, req).Return(nil, nil, errExpected).Once()

	// act
	res, err := ucts.controller.DeleteRecords(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected.Error(), rpcStatus.Message())
	ucts.assertAllMocks()
}
