package operations_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"truenorth/services/operations_service/repositories/operations"
)

type OperationRepoTestSuite struct {
	suite.Suite

	repo      *operations.OperationsRepoImpl
	sqlDbMock *sql.DB
	gormDb    *gorm.DB
	sqlMock   sqlmock.Sqlmock
	ctx       context.Context
}

func (ucts *OperationRepoTestSuite) SetupTest() {
	// Initialize necessary dependencies
	ucts.sqlDbMock, ucts.sqlMock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       ucts.sqlDbMock,
		DriverName: "postgres",
	})
	ucts.gormDb, _ = gorm.Open(dialector, &gorm.Config{})
	ucts.ctx = context.TODO()

	ucts.repo = &operations.OperationsRepoImpl{}
	ucts.repo.SetDbInstance(ucts.gormDb)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &OperationRepoTestSuite{})
}
