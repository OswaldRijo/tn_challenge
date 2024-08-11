package users_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	apimocks "truenorth/mocks/users_service/api"
	"truenorth/services/users_service/controllers/users"
)

type UserControllerTestSuite struct {
	suite.Suite

	controller     *users.UsersControllerImpl
	ctx            context.Context
	usersApiMock   *apimocks.UserApi
	assertAllMocks func()
}

func (ucts *UserControllerTestSuite) SetupTest() {

	// Initialize necessary dependencies
	ucts.usersApiMock = new(apimocks.UserApi)
	ucts.ctx = context.TODO()

	// Define a function to assert expectations on mocks
	ucts.assertAllMocks = func() {
		ucts.usersApiMock.AssertExpectations(ucts.T())
	}

	ucts.controller = &users.UsersControllerImpl{}
	ucts.controller.SetUserApi(ucts.usersApiMock)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &UserControllerTestSuite{})
}
