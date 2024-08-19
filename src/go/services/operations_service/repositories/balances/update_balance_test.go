package balances_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	testdata "truenorth/services/operations_service/test"
)

func (bri *BalanceRepoTestSuite) Test_UpdateBalance_Success() {
	// arrange
	bri.sqlMock.ExpectBegin()
	bri.sqlMock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "balances" SET "created_at"=$1,"updated_at"=$2,"user_id"=$3,"current_balance"=$4 WHERE "id" = $5`)).
		WithArgs(testdata.YagamiLightBalanceModel.CreatedAt, sqlmock.AnyArg(), testdata.YagamiLightBalanceModel.UserID, testdata.YagamiLightBalanceModel.CurrentBalance, testdata.YagamiLightBalanceModel.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	tx := bri.gormDb.Begin()

	// act
	err := bri.repo.UpdateBalance(bri.ctx, testdata.YagamiLightBalanceModel, tx)

	// assert
	bri.Nil(err)
}

func (bri *BalanceRepoTestSuite) Test_UpdateBalance_ErrorCreatingOperations() {
	// arrange
	expectedErr := errors.New("something went wrong")
	bri.sqlMock.ExpectBegin()
	bri.sqlMock.
		ExpectExec(regexp.QuoteMeta(`UPDATE "balances" SET "created_at"=$1,"updated_at"=$2,"user_id"=$3,"current_balance"=$4 WHERE "id" = $5`)).
		WithArgs(testdata.YagamiLightBalanceModel.CreatedAt, sqlmock.AnyArg(), testdata.YagamiLightBalanceModel.UserID, testdata.YagamiLightBalanceModel.CurrentBalance, testdata.YagamiLightBalanceModel.ID).WillReturnError(expectedErr)
	tx := bri.gormDb.Begin()
	// act
	err := bri.repo.UpdateBalance(bri.ctx, testdata.YagamiLightBalanceModel, tx)

	// assert
	bri.NotNil(err)
	bri.Equal(expectedErr, err)
}
