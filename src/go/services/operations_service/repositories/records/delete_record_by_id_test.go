package records_test

import (
	"database/sql/driver"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	operationsmodels "truenorth/services/operations_service/models"
)

func (rrt *RecordsRepoTestSuite) Test_DeleteRecordById_Success() {
	// arrange
	recordId := int64(1)
	newRecord := &operationsmodels.Record{
		ID:                recordId,
		OperationID:       1,
		UserID:            1,
		UserBalance:       999999,
		Deleted:           false,
		OperationResponse: "3",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	rrt.sqlMock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "operation_id", "user_id", "user_balance", "deleted", "operation_response", "created_at", "updated_at"}).
		AddRow(newRecord.ID, newRecord.OperationID, newRecord.UserID, newRecord.UserBalance, newRecord.Deleted, newRecord.OperationResponse, newRecord.CreatedAt, newRecord.UpdatedAt)

	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "records" WHERE "deleted" = $1 AND "id" = $2 ORDER BY "records"."id" LIMIT $3`)).
		WithArgs(newRecord.Deleted, newRecord.ID, 1).WillReturnRows(rows)

	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "operations" WHERE "operations"."id" = $1`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "operation_id", "user_id", "user_balance", "deleted", "operation_response", "created_at", "updated_at"}))

	rrt.sqlMock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "records" SET "created_at"=$1,"updated_at"=$2,"operation_id"=$3,"user_id"=$4,"user_balance"=$5,"deleted"=$6,"operation_response"=$7 WHERE "id" = $8`)).
		WithArgs(newRecord.CreatedAt, sqlmock.AnyArg(), newRecord.OperationID, newRecord.UserID, newRecord.UserBalance, true, newRecord.OperationResponse, newRecord.ID).WillReturnResult(driver.ResultNoRows)

	rrt.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "records" ("created_at","updated_at","operation_id","user_id","user_balance","deleted","operation_response","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT ("id") DO UPDATE SET "updated_at"=$9,"operation_id"="excluded"."operation_id","user_id"="excluded"."user_id","user_balance"="excluded"."user_balance","deleted"="excluded"."deleted","operation_response"="excluded"."operation_response" RETURNING "id"`)).
		WithArgs(newRecord.CreatedAt, sqlmock.AnyArg(), newRecord.OperationID, newRecord.UserID, newRecord.UserBalance, true, newRecord.OperationResponse, newRecord.ID, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newRecord.ID))
	tx := rrt.gormDb.Begin()

	// act
	err := rrt.repo.DeleteRecordById(rrt.ctx, recordId, tx)

	// assert
	rrt.Nil(err)
}
