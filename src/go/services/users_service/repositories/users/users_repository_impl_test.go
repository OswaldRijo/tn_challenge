package users_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"truenorth/services/users_service/repositories/users"
)

type UserRepositoryTestSuite struct {
	suite.Suite

	repository     *users.UsersRepoImpl
	sqlDbMock      *sql.DB
	gormDb         *gorm.DB
	sqlMock        sqlmock.Sqlmock
	ctx            context.Context
	assertAllMocks func()
}

func (ucts *UserRepositoryTestSuite) SetupTest() {
	// Initialize necessary dependencies
	ucts.sqlDbMock, ucts.sqlMock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       ucts.sqlDbMock,
		DriverName: "postgres",
	})
	ucts.gormDb, _ = gorm.Open(dialector, &gorm.Config{})
	ucts.ctx = context.TODO()

	ucts.repository = &users.UsersRepoImpl{}
	ucts.repository.SetDbInstance(ucts.gormDb)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &UserRepositoryTestSuite{})
}
