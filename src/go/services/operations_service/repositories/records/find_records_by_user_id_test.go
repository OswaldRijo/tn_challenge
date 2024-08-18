package records_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	testdata "truenorth/services/users_service/test"
)

func (rrt *RecordsRepoTestSuite) Test_FindRecordsByUserId_Success() {
	// arrange
	expectedId := int64(1)
	userId := int64(1)
	limit := int32(1)
	page := int32(1)
	count := int64(1)
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "cost", "operation_type", "args"}).
		AddRow(expectedId, testdata.OperationModel.CreatedAt, testdata.OperationModel.UpdatedAt, testdata.OperationModel.UserID, testdata.OperationModel.Cost, testdata.OperationModel.OperationType, testdata.OperationModel.Args)
	countRow := sqlmock.NewRows([]string{"count"}).
		AddRow(count)
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "records" WHERE "deleted" = $1 AND "user_id" = $2 LIMIT $3 OFFSET $4`)).
		WithArgs(false, userId, limit, page).WillReturnRows(rows)

	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "records" WHERE "deleted" = $1 AND "user_id" = $2`)).
		WithArgs(false, userId).WillReturnRows(countRow)

	// act
	records, totalCount, err := rrt.repo.FindRecordsByUserId(rrt.ctx, userId, limit, page)

	// assert
	rrt.Nil(err)
	rrt.Equal(count, totalCount)
	rrt.Len(records, 1)
}

func (rrt *RecordsRepoTestSuite) Test_FindRecordsByUserId_ErrorGettingRecords() {
	// arrange
	userId := int64(1)
	limit := int32(1)
	page := int32(1)
	zeroCount := int64(0)
	expectedError := errors.New("oops")
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "records" WHERE "deleted" = $1 AND "user_id" = $2 LIMIT $3 OFFSET $4`)).
		WithArgs(false, userId, limit, page).WillReturnError(expectedError)

	// act
	records, totalCount, err := rrt.repo.FindRecordsByUserId(rrt.ctx, userId, limit, page)

	// assert
	rrt.Nil(records)
	rrt.Equal(zeroCount, totalCount)
	rrt.NotNil(err)
	rrt.Equal(expectedError.Error(), err.Error())
}

func (rrt *RecordsRepoTestSuite) Test_FindRecordsByUserId_ErrorGettingCount() {
	// arrange
	expectedId := int64(1)
	userId := int64(1)
	limit := int32(1)
	page := int32(1)
	zeroCount := int64(0)
	expectedError := errors.New("oops")
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "cost", "operation_type", "args"}).
		AddRow(expectedId, testdata.OperationModel.CreatedAt, testdata.OperationModel.UpdatedAt, testdata.OperationModel.UserID, testdata.OperationModel.Cost, testdata.OperationModel.OperationType, testdata.OperationModel.Args)
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "records" WHERE "deleted" = $1 AND "user_id" = $2 LIMIT $3 OFFSET $4`)).
		WithArgs(false, userId, limit, page).WillReturnRows(rows)

	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "records" WHERE "deleted" = $1 AND "user_id" = $2`)).
		WithArgs(false, userId).WillReturnError(expectedError)

	// act
	records, totalCount, err := rrt.repo.FindRecordsByUserId(rrt.ctx, userId, limit, page)

	// assert
	rrt.Nil(records)
	rrt.Equal(zeroCount, totalCount)
	rrt.NotNil(err)
	rrt.Equal(expectedError.Error(), err.Error())
}
