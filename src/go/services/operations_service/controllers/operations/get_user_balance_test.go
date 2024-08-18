package operations_test

import (
	"errors"

	"google.golang.org/grpc/status"

	operationspb "truenorth/pb/operations"
	"truenorth/services/users_service/test"
)

func (ucts *OperationControllerTestSuite) Test_GetUserBalance_Success() {
	// arrange
	req := &operationspb.GetUserBalanceRequest{UserId: 1}
	ucts.operationsApiMock.On("GetUserBalance", ucts.ctx, req.GetUserId()).Return(test.YagamiLightBalancePb, nil).Once()

	// act
	res, err := ucts.controller.GetUserBalance(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.YagamiLightBalancePb.CurrentBalance, res.GetBalance().GetCurrentBalance())
	ucts.Equal(test.YagamiLightBalancePb.UserId, res.GetBalance().GetUserId())
	ucts.assertAllMocks()
}

func (ucts *OperationControllerTestSuite) Test_GetUserBalance_ApiError() {
	// arrange
	req := &operationspb.GetUserBalanceRequest{UserId: 1}
	errExpected := errors.New("error getting user balance")
	ucts.operationsApiMock.On("GetUserBalance", ucts.ctx, req.GetUserId()).Return(nil, errExpected).Once()

	// act
	res, err := ucts.controller.GetUserBalance(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected.Error(), rpcStatus.Message())
	ucts.assertAllMocks()
}
