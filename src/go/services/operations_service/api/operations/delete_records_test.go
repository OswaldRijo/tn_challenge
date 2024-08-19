package operations_test

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/api/operations"
	"truenorth/services/operations_service/config"
	operationsmodels "truenorth/services/operations_service/models"
	"truenorth/services/operations_service/test"
)

func (ucts *OperationApiTestSuite) Test_DeleteUser_Success() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(test.RecordsModel, nil).Once()
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, test.YagamiLightBalanceModel.UserID).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.recordsRepoMock.On("DeleteRecordById", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.balancesRepoMock.On("UpdateBalance", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Len(records, len(test.RecordsPb))
	ucts.Equal(test.RecordsPb[0].OperationResponse, records[0].OperationResponse)
	ucts.Equal(test.YagamiLightBalancePb.GetUserId(), balance.GetUserId())
	ucts.Equal(test.YagamiLightBalancePb.GetCurrentBalance(), balance.GetCurrentBalance())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorRecordsNotFound() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := gorm.ErrRecordNotFound
	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(nil, expectedErr).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorFindingRecords() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("some error")
	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(nil, expectedErr).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorGetBalanceByUserId() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("some error")
	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(test.RecordsModel, nil).Once()
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, test.YagamiLightBalanceModel.UserID).Return(nil, expectedErr).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorRecordDoesNotBelongsToUser() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New(operations.RecordDoesNotBelongToTheUser)
	recordId := int64(1)
	userId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: userId}
	recordArr := *test.RecordsModel[0]
	(&recordArr).UserID = 1
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return([]*operationsmodels.Record{&recordArr}, nil).Once()
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, userId).Return(test.YagamiLightBalanceModel, nil).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorDeletingRecordNotFound() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := gorm.ErrRecordNotFound
	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(test.RecordsModel, nil).Once()
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, test.YagamiLightBalanceModel.UserID).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.recordsRepoMock.On("DeleteRecordById", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorDeletingRecord() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("error deleting balance")
	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(test.RecordsModel, nil).Once()
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, test.YagamiLightBalanceModel.UserID).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.recordsRepoMock.On("DeleteRecordById", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_DeleteUser_ErrorUpdatingBalance() {
	// arrange
	config.Config.DefaultUserBalance = 500
	operations.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}
	expectedErr := errors.New("error updating balance")
	recordId := int64(1)
	req := &operationspb.DeleteRecordsRequest{RecordIds: []int64{recordId}, UserId: test.YagamiLightBalanceModel.UserID}
	ucts.recordsRepoMock.On("FindRecordsByIds", ucts.ctx, recordId).Return(test.RecordsModel, nil).Once()
	ucts.balancesRepoMock.On("GetBalanceByUserId", ucts.ctx, test.YagamiLightBalanceModel.UserID).Return(test.YagamiLightBalanceModel, nil).Once()
	ucts.recordsRepoMock.On("DeleteRecordById", ucts.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	ucts.balancesRepoMock.On("UpdateBalance", ucts.ctx, mock.Anything, mock.Anything).Return(expectedErr).Once()

	// act
	records, balance, err := ucts.api.DeleteRecord(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.Nil(balance)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}
