package operations_test

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/controllers/operations"
	"truenorth/services/users_service/test"
)

func (ucts *OperationControllerTestSuite) Test_ApplyOperation_Success() {
	// arrange
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.operationsApiMock.On("ApplyOperation", ucts.ctx, req).Return(test.OperationPb, test.RecordPb, test.YagamiLightBalancePb, nil).Once()

	// act
	res, err := ucts.controller.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.RecordPb, res.GetRecord())
	ucts.Equal(test.YagamiLightBalancePb, res.GetCurrentUserBalance())
	ucts.Equal(test.OperationPb, res.GetOperation())
	ucts.assertAllMocks()
}

func (ucts *OperationControllerTestSuite) Test_ApplyOperation_WrongOperation() {
	// arrange
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: 8, Args: []float64{1, 2}, UserId: userId}

	// act
	res, err := ucts.controller.ApplyOperation(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(codes.InvalidArgument, rpcStatus.Code())
	ucts.Equal(operations.InvalidOperationNumber, rpcStatus.Message())
	ucts.assertAllMocks()
}

func (ucts *OperationControllerTestSuite) Test_ApplyOperation_ApiError() {
	// arrange
	userId := int64(1)
	errExpected := errors.New("error getting user balance")
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.operationsApiMock.On("ApplyOperation", ucts.ctx, req).Return(nil, nil, nil, errExpected).Once()

	// act
	res, err := ucts.controller.ApplyOperation(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected.Error(), rpcStatus.Message())
	ucts.assertAllMocks()
}
