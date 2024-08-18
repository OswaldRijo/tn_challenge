package operations_test

import (
	"errors"

	"truenorth/services/operations_service/api/operations"
	"truenorth/services/users_service/test"
)

func (ucts *OperationApiTestSuite) Test_GetUserBalance_Success() {
	// arrange
	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(test.YagamiLightBalanceModel, nil).Once()

	// act
	balance, err := ucts.api.GetUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.YagamiLightBalancePb.GetUserId(), balance.GetUserId())
	ucts.Equal(test.YagamiLightBalancePb.GetCurrentBalance(), balance.GetCurrentBalance())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_GetUserBalance_ErrorGettingBalance() {
	// arrange
	expectedError := errors.New("some error")
	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(nil, expectedError).Once()

	// act
	balance, err := ucts.api.GetUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedError.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_GetUserBalance_ErrorBalanceNotFound() {
	// arrange
	expectedError := errors.New(operations.UserBalanceNotFound)
	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(nil, nil).Once()

	// act
	balance, err := ucts.api.GetUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedError.Error(), err.Error())
	ucts.assertAllMocks()
}
