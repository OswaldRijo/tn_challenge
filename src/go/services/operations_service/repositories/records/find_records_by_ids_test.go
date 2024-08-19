package records_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	testdata "truenorth/services/operations_service/test"
)

func (rrt *RecordsRepoTestSuite) Test_FindRecordsByIds_Success() {
	// arrange
	expectedId := int64(1)
	recordsId := int64(1)
	rows := sqlmock.NewRows([]string{"id", "operation_id", "user_id", "user_balance", "deleted", "operation_response", "created_at", "updated_at"}).
		AddRow(expectedId, testdata.RecordModel.OperationID, testdata.RecordModel.UserID, testdata.RecordModel.UserBalance, testdata.RecordModel.Deleted, testdata.RecordModel.OperationResponse, testdata.RecordModel.CreatedAt, testdata.RecordModel.UpdatedAt)
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "records" WHERE "id" = $1`)).
		WithArgs(recordsId).WillReturnRows(rows)
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "operations" WHERE "operations"."id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "operation_id", "user_id", "user_balance", "deleted", "operation_response", "created_at", "updated_at"}))

	// act
	records, err := rrt.repo.FindRecordsByIds(rrt.ctx, recordsId)

	// assert
	rrt.Nil(err)
	rrt.Len(records, 1)
}

func (rrt *RecordsRepoTestSuite) Test_FindRecordsByIds_ErrorFiltering() {
	// arrange
	recordsId := int64(1)
	expectedError := errors.New("filtering failed")
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "records" WHERE "id" = $1`)).
		WithArgs(recordsId).WillReturnError(expectedError)

	// act
	operations, err := rrt.repo.FindRecordsByIds(rrt.ctx, recordsId)

	// assert
	rrt.Nil(operations)
	rrt.NotNil(err)
	rrt.Equal(expectedError.Error(), err.Error())
}
