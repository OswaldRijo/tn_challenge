package records_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"truenorth/services/operations_service/repositories/records"
)

type RecordsRepoTestSuite struct {
	suite.Suite

	repo      *records.RecordsRepoImpl
	sqlDbMock *sql.DB
	gormDb    *gorm.DB
	sqlMock   sqlmock.Sqlmock
	ctx       context.Context
}

func (rrt *RecordsRepoTestSuite) SetupTest() {
	// Initialize necessary dependencies
	rrt.sqlDbMock, rrt.sqlMock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       rrt.sqlDbMock,
		DriverName: "postgres",
	})
	rrt.gormDb, _ = gorm.Open(dialector, &gorm.Config{})
	rrt.ctx = context.TODO()

	rrt.repo = &records.RecordsRepoImpl{}
	rrt.repo.SetDbInstance(rrt.gormDb)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &RecordsRepoTestSuite{})
}
