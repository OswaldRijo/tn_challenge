package balances_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	testdata "truenorth/services/operations_service/test"
)

func (bri *BalanceRepoTestSuite) Test_CreateBalance_Success() {
	// arrange
	expectedId := int64(1)
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(expectedId)
	bri.sqlMock.ExpectBegin()
	bri.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "balances" ("created_at","updated_at","user_id","current_balance","id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(testdata.YagamiLightBalanceModel.CreatedAt, testdata.YagamiLightBalanceModel.UpdatedAt, testdata.YagamiLightBalanceModel.UserID, testdata.YagamiLightBalanceModel.CurrentBalance, testdata.YagamiLightBalanceModel.ID).WillReturnRows(rows)
	tx := bri.gormDb.Begin()

	// act
	err := bri.repo.CreateBalance(bri.ctx, testdata.YagamiLightBalanceModel, tx)

	// assert
	bri.Nil(err)
}

func (bri *BalanceRepoTestSuite) Test_CreateBalance_ErrorCreatingOperations() {
	// arrange
	expectedErr := errors.New("something went wrong")
	bri.sqlMock.ExpectBegin()
	bri.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "balances" ("created_at","updated_at","user_id","current_balance","id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(testdata.YagamiLightBalanceModel.CreatedAt, testdata.YagamiLightBalanceModel.UpdatedAt, testdata.YagamiLightBalanceModel.UserID, testdata.YagamiLightBalanceModel.CurrentBalance, testdata.YagamiLightBalanceModel.ID).WillReturnError(expectedErr)
	tx := bri.gormDb.Begin()
	// act
	err := bri.repo.CreateBalance(bri.ctx, testdata.YagamiLightBalanceModel, tx)

	// assert
	bri.NotNil(err)
	bri.Equal(expectedErr, err)
}
