package operations_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	testdata "truenorth/services/users_service/test"
)

func (ucts *OperationRepoTestSuite) Test_CreateOperation_Success() {
	// arrange
	expectedId := int64(1)
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(expectedId)
	ucts.sqlMock.ExpectBegin()
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "operations" ("created_at","updated_at","user_id","cost","operation_type","args","id") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(testdata.OperationModel.CreatedAt, testdata.OperationModel.UpdatedAt, testdata.OperationModel.UserID, testdata.OperationModel.Cost, testdata.OperationModel.OperationType, testdata.OperationModel.Args, testdata.OperationModel.ID).WillReturnRows(rows)
	tx := ucts.gormDb.Begin()
	// act
	err := ucts.repo.CreateOperation(ucts.ctx, testdata.OperationModel, tx)

	// assert
	ucts.Nil(err)
}

func (ucts *OperationRepoTestSuite) Test_CreateOperation_ErrorCreatingOperations() {
	// arrange
	expectedErr := errors.New("something went wrong")
	ucts.sqlMock.ExpectBegin()
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "operations" ("created_at","updated_at","user_id","cost","operation_type","args","id") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(testdata.OperationModel.CreatedAt, testdata.OperationModel.UpdatedAt, testdata.OperationModel.UserID, testdata.OperationModel.Cost, testdata.OperationModel.OperationType, testdata.OperationModel.Args, testdata.OperationModel.ID).WillReturnError(expectedErr)
	tx := ucts.gormDb.Begin()
	// act
	err := ucts.repo.CreateOperation(ucts.ctx, testdata.OperationModel, tx)

	// assert
	ucts.NotNil(err)
	ucts.Equal(expectedErr, err)
}
