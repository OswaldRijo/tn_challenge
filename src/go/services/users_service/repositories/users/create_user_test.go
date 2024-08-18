package users_test

import (
	"fmt"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"

	"truenorth/services/users_service/models"
	"truenorth/services/users_service/test"
)

func (ucts *UserRepositoryTestSuite) Test_CreateUser_Success() {
	// arrange
	expectedId := int64(10)
	rows := sqlmock.NewRows([]string{"ID"}).AddRow(expectedId)
	ucts.sqlMock.ExpectBegin()
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","password","status","created_at","updated_at") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(test.TheGoat.Username, test.TheGoat.Password, models.StatusIntToNameMap[models.StatusActive], test.TheGoat.CreatedAt, test.TheGoat.UpdatedAt).WillReturnRows(rows)
	ucts.sqlMock.ExpectCommit()
	tx := ucts.gormDb.Begin()

	// act
	err := ucts.repository.CreateUser(ucts.ctx, test.TheGoat, tx)

	// assert
	ucts.Nil(err)
	ucts.Equal(expectedId, test.TheGoat.ID)
}

func (ucts *UserRepositoryTestSuite) Test_CreateUser_Error() {
	// arrange
	expectedErr := fmt.Errorf("error inserting user")
	ucts.sqlMock.ExpectBegin()
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("username","password","status","created_at","updated_at") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(test.TheGoat.Username, test.TheGoat.Password, models.StatusIntToNameMap[models.StatusActive], test.TheGoat.CreatedAt, test.TheGoat.UpdatedAt).WillReturnError(expectedErr)
	ucts.sqlMock.ExpectCommit()
	tx := ucts.gormDb.Begin()

	// act
	err := ucts.repository.CreateUser(ucts.ctx, test.TheGoat, tx)

	// assert
	ucts.Equal(expectedErr, err)
}
