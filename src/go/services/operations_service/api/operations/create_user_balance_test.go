package operations_test

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"truenorth/services/operations_service/api/operations"
	"truenorth/services/operations_service/config"
	"truenorth/services/operations_service/test"
)

func (ucts *OperationApiTestSuite) Test_CreateUserBalance_Success() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(nil, nil).Once()
	ucts.balancesRepoMock.On("CreateBalance", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	// act
	balance, err := ucts.api.CreateUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(err)
	ucts.Equal(userId, balance.GetUserId())
	ucts.Equal(config.Config.DefaultUserBalance, balance.GetCurrentBalance())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_CreateUserBalance_AlreadyExists() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(test.YagamiLightBalanceModel, nil).Once()

	// act
	balance, err := ucts.api.CreateUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.YagamiLightBalancePb.GetUserId(), balance.GetUserId())
	ucts.Equal(test.YagamiLightBalancePb.GetCurrentBalance(), balance.GetCurrentBalance())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_CreateUserBalance_ErrorCreatingUserBalance() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("error creating user balance")
	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(nil, nil).Once()
	ucts.balancesRepoMock.On("CreateBalance", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	balance, err := ucts.api.CreateUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_CreateUserBalance_ErrorGettingUserBalance() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("error creating user balance")
	userId := int64(1)
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(nil, expectedErr).Once()

	// act
	balance, err := ucts.api.CreateUserBalance(ucts.ctx, userId)

	// assert
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}
