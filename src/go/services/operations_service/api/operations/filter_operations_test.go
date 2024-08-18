package operations_test

import (
	"errors"

	operationspb "truenorth/pb/operations"
	"truenorth/services/users_service/test"
)

func (ucts *OperationApiTestSuite) Test_FindRecordsByUserId_Success() {
	// arrange
	expectedTotalCount := int64(1)
	userId := int64(1)
	page := int32(1)
	limit := int32(1)
	req := &operationspb.FilterRecordsRequest{UserId: &userId, Limit: &limit, Page: &page}
	ucts.recordsRepoMock.On("FindRecordsByUserId", ucts.ctx, userId, limit, page).Return(test.RecordsModel, expectedTotalCount, nil).Once()

	// act
	records, totalCount, err := ucts.api.FilterRecords(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Len(records, len(test.RecordsPb))
	ucts.Equal(test.RecordsPb[0].OperationResponse, records[0].OperationResponse)
	ucts.Equal(expectedTotalCount, totalCount)
	ucts.assertAllMocks()
}

func (ucts *OperationApiTestSuite) Test_FindRecordsByUserId_ErrorFilteringRecords() {
	// arrange
	expectedError := errors.New("some error")
	expectedTotalCount := int64(0)
	userId := int64(1)
	page := int32(1)
	limit := int32(1)
	req := &operationspb.FilterRecordsRequest{UserId: &userId, Limit: &limit, Page: &page}
	ucts.recordsRepoMock.On("FindRecordsByUserId", ucts.ctx, userId, limit, page).Return(nil, expectedTotalCount, expectedError).Once()

	// act
	records, totalCount, err := ucts.api.FilterRecords(ucts.ctx, req)

	// assert
	ucts.Nil(records)
	ucts.NotNil(err)
	ucts.Equal(expectedError.Error(), err.Error())
	ucts.Equal(expectedTotalCount, totalCount)
	ucts.assertAllMocks()
}
