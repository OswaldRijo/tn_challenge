package balances_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"truenorth/services/operations_service/repositories/balances"
)

type BalanceRepoTestSuite struct {
	suite.Suite

	repo      *balances.BalancesRepoImpl
	sqlDbMock *sql.DB
	gormDb    *gorm.DB
	sqlMock   sqlmock.Sqlmock
	ctx       context.Context
}

func (bri *BalanceRepoTestSuite) SetupTest() {
	// Initialize necessary dependencies
	bri.sqlDbMock, bri.sqlMock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       bri.sqlDbMock,
		DriverName: "postgres",
	})
	bri.gormDb, _ = gorm.Open(dialector, &gorm.Config{})
	bri.ctx = context.TODO()

	bri.repo = &balances.BalancesRepoImpl{}
	bri.repo.SetDbInstance(bri.gormDb)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &BalanceRepoTestSuite{})
}
