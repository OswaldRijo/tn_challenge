package users_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	repomocks "truenorth/mocks/users_service/reposirories"
	"truenorth/services/users_service/api/users"
)

type UserApiTestSuite struct {
	suite.Suite

	apiUser               *users.UsersApiImpl
	ctx                   context.Context
	usersRepositoriesMock *repomocks.UsersRepo
	assertAllMocks        func()
}

func (ucts *UserApiTestSuite) SetupTest() {
	// Initialize necessary dependencies
	ucts.usersRepositoriesMock = new(repomocks.UsersRepo)
	ucts.ctx = context.TODO()

	// Define a function to assert expectations on mocks
	ucts.assertAllMocks = func() {
		ucts.usersRepositoriesMock.AssertExpectations(ucts.T())
	}

	ucts.apiUser = &users.UsersApiImpl{}
	ucts.apiUser.SetUserMock(ucts.usersRepositoriesMock)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &UserApiTestSuite{})
}
