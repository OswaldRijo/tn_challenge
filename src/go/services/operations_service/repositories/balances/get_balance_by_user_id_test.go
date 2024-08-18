package balances_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	testdata "truenorth/services/users_service/test"
)

func (bri *BalanceRepoTestSuite) Test_GetBalanceByUserId_Success() {
	// arrange
	rows := sqlmock.NewRows([]string{"id", "user_id", "current_balance", "created_at", "updated_at"}).
		AddRow(testdata.YagamiLightBalanceModel.ID, testdata.YagamiLightBalanceModel.UserID, testdata.YagamiLightBalanceModel.CurrentBalance, testdata.YagamiLightBalanceModel.CreatedAt, testdata.YagamiLightBalanceModel.UpdatedAt)
	bri.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "balances" WHERE "user_id" = $1 ORDER BY "balances"."id" LIMIT $2`)).
		WithArgs(testdata.YagamiLightBalanceModel.UserID, sqlmock.AnyArg()).WillReturnRows(rows)

	// act
	balance, err := bri.repo.GetBalanceByUserId(bri.ctx, testdata.YagamiLightBalanceModel.UserID)

	// assert
	bri.Nil(err)
	bri.Equal(testdata.YagamiLightBalanceModel, balance)
}

func (bri *BalanceRepoTestSuite) Test_GetBalanceByUserId_BalanceNotFound() {
	// arrange
	expectedErr := gorm.ErrRecordNotFound
	bri.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "balances" WHERE "user_id" = $1 ORDER BY "balances"."id" LIMIT $2`)).
		WithArgs(testdata.YagamiLightBalanceModel.UserID, sqlmock.AnyArg()).WillReturnError(expectedErr)
	// act
	balance, err := bri.repo.GetBalanceByUserId(bri.ctx, testdata.YagamiLightBalanceModel.UserID)

	// assert
	bri.Nil(balance)
	bri.Nil(err)
}

func (bri *BalanceRepoTestSuite) Test_GetBalanceByUserId_ErrorGettingBalance() {
	// arrange
	expectedErr := errors.New("something went wrong")
	bri.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "balances" WHERE "user_id" = $1 ORDER BY "balances"."id" LIMIT $2`)).
		WithArgs(testdata.YagamiLightBalanceModel.UserID, sqlmock.AnyArg()).WillReturnError(expectedErr)
	// act
	balance, err := bri.repo.GetBalanceByUserId(bri.ctx, testdata.YagamiLightBalanceModel.UserID)

	// assert
	bri.Nil(balance)
	bri.NotNil(err)
	bri.Equal(expectedErr, err)
}
