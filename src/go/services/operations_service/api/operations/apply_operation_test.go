package operations_test

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/api/operations"
	"truenorth/services/operations_service/test"
)

func (ucts *OperationApiTestSuite) Test_ApplyOperation_Success() {
	// arrange
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, req.GetUserId()).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.operationsRepoMock.On("CreateOperation", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.recordsRepoMock.On("CreateRecord", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.balancesRepoMock.On("UpdateBalance", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	// act
	op, record, balance, err := ucts.api.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(test.RecordPb.GetUserBalance(), record.GetUserBalance())
	ucts.Equal(test.RecordPb.GetDeleted(), record.GetDeleted())
	ucts.Equal(test.YagamiLightBalancePb.GetUserId(), balance.GetUserId())
	ucts.Equal(test.YagamiLightBalancePb.GetCurrentBalance(), balance.GetCurrentBalance())
	ucts.Equal(test.OperationPb.GetUserId(), op.GetUserId())
	ucts.Equal(test.OperationPb.GetOperationType(), op.GetOperationType())
	ucts.Equal(test.OperationPb.GetCost(), op.GetCost())
	ucts.Equal(test.OperationPb.GetArgs(), op.GetArgs())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_ApplyOperation_ApiErrorGettingBalance() {
	// arrange
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("error getting balance")
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, req.GetUserId()).Return(nil, expectedErr).Once()

	// act
	op, record, balance, err := ucts.api.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(op)
	ucts.Nil(record)
	ucts.Nil(balance)

	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_ApplyOperation_ApiErrorBalanceNotFound() {
	// arrange
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New(operations.UserBalanceNotFound)
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, req.GetUserId()).Return(nil, nil).Once()

	// act
	op, record, balance, err := ucts.api.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(op)
	ucts.Nil(record)
	ucts.Nil(balance)

	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_ApplyOperation_ApiErrorInsertingModel() {
	// arrange
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("inserting operation error")
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, req.GetUserId()).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.operationsRepoMock.On("CreateOperation", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	op, record, balance, err := ucts.api.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(op)
	ucts.Nil(record)
	ucts.Nil(balance)

	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_ApplyOperation_ApiErrorInsertingRecord() {
	// arrange
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("inserting record error")
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, req.GetUserId()).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.operationsRepoMock.On("CreateOperation", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.recordsRepoMock.On("CreateRecord", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	op, record, balance, err := ucts.api.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(op)
	ucts.Nil(record)
	ucts.Nil(balance)

	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_ApplyOperation_ApiErrorUpdatingBalance() {
	// arrange
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("update balance error")
	userId := int64(1)
	req := &operationspb.ApplyOperationRequest{OperationType: operationspb.OperationType_ADDITION, Args: []float64{1, 2}, UserId: userId}
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, req.GetUserId()).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.operationsRepoMock.On("CreateOperation", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.recordsRepoMock.On("CreateRecord", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.balancesRepoMock.On("UpdateBalance", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	op, record, balance, err := ucts.api.ApplyOperation(ucts.ctx, req)

	// assert
	ucts.Nil(op)
	ucts.Nil(record)
	ucts.Nil(balance)

	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}
