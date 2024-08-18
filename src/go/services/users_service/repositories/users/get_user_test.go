package users_test

import (
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"truenorth/services/users_service/models"
	testdata "truenorth/services/users_service/test"
)

func (ucts *UserRepositoryTestSuite) Test_GetUser_Success() {
	// arrange
	expectedId := int64(10)
	expectedLimit := int64(1)
	rows := sqlmock.NewRows([]string{"id", "username", "password", "status", "created_at", "updated_at"}).AddRow(expectedId, testdata.TheGoat.Username, testdata.TheGoat.Password, models.StatusIntToNameMap[models.StatusActive], testdata.TheGoat.CreatedAt, testdata.TheGoat.UpdatedAt)
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "userId" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(expectedId, expectedLimit).WillReturnRows(rows)

	// act
	user, err := ucts.repository.GetUser(ucts.ctx, map[string]interface{}{"userId": expectedId})

	// assert
	ucts.Nil(err)
	ucts.Equal(testdata.TheGoat.Username, user.Username)
	ucts.Equal(testdata.TheGoat.Password, user.Password)
	ucts.Equal(testdata.TheGoat.Status, user.Status)
	ucts.Equal(expectedId, user.ID)
}

func (ucts *UserRepositoryTestSuite) Test_GetUser_Error() {
	// arrange
	expectedId := int64(10)
	expectedLimit := int64(1)
	expectedError := errors.New("error getting user")
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "userId" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(expectedId, expectedLimit).WillReturnError(expectedError)

	// act
	user, err := ucts.repository.GetUser(ucts.ctx, map[string]interface{}{"userId": expectedId})

	// assert
	ucts.NotNil(err)
	ucts.Nil(user)
}

func (ucts *UserRepositoryTestSuite) Test_GetUser_ErrorNotFound() {
	// arrange
	expectedId := int64(10)
	expectedLimit := int64(1)
	expectedError := gorm.ErrRecordNotFound
	ucts.sqlMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "userId" = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(expectedId, expectedLimit).WillReturnError(expectedError)

	// act
	user, err := ucts.repository.GetUser(ucts.ctx, map[string]interface{}{"userId": expectedId})

	// assert
	ucts.Nil(err)
	ucts.Nil(user)
}
