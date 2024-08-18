package records_test

import (
	"errors"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	operationsmodels "truenorth/services/operations_service/models"
)

func (rrt *RecordsRepoTestSuite) Test_CreateRecord_Success() {
	// arrange
	expectedId := int64(1)
	newRecord := &operationsmodels.Record{
		OperationID:       1,
		UserID:            1,
		UserBalance:       999999,
		Deleted:           false,
		OperationResponse: "3",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(expectedId)
	rrt.sqlMock.ExpectBegin()
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "records" ("created_at","updated_at","operation_id","user_id","user_balance","deleted","operation_response") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(newRecord.CreatedAt, newRecord.UpdatedAt, newRecord.OperationID, newRecord.UserID, newRecord.UserBalance, newRecord.Deleted, newRecord.OperationResponse).WillReturnRows(rows)
	tx := rrt.gormDb.Begin()

	// act
	err := rrt.repo.CreateRecord(rrt.ctx, newRecord, tx)

	// assert
	rrt.Nil(err)
	rrt.Equal(expectedId, newRecord.ID)
}

func (rrt *RecordsRepoTestSuite) Test_CreateRecord_ErrorCreating() {
	// arrange
	expectedErr := errors.New("error creating record")
	newRecord := &operationsmodels.Record{
		OperationID:       1,
		UserID:            1,
		UserBalance:       999999,
		Deleted:           false,
		OperationResponse: "3",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	rrt.sqlMock.ExpectBegin()
	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "records" ("created_at","updated_at","operation_id","user_id","user_balance","deleted","operation_response") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(newRecord.CreatedAt, newRecord.UpdatedAt, newRecord.OperationID, newRecord.UserID, newRecord.UserBalance, newRecord.Deleted, newRecord.OperationResponse).WillReturnError(expectedErr)
	tx := rrt.gormDb.Begin()

	// act
	err := rrt.repo.CreateRecord(rrt.ctx, newRecord, tx)

	// assert
	rrt.Equal(int64(0), newRecord.ID)
	rrt.NotNil(err)
	rrt.Equal(expectedErr.Error(), err.Error())
}
