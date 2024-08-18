package users_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	pubsubmock "truenorth/mocks/packages/pubsub"
	repomocks "truenorth/mocks/users_service/reposirories"
	"truenorth/services/users_service/api/users"
	"truenorth/services/users_service/config"
)

type UserApiTestSuite struct {
	suite.Suite

	apiUser               *users.UsersApiImpl
	ctx                   context.Context
	producerMock          *pubsubmock.Producer
	usersRepositoriesMock *repomocks.UsersRepo
	assertAllMocks        func()
}

func (ucts *UserApiTestSuite) SetupTest() {
	// Initialize necessary dependencies
	ucts.usersRepositoriesMock = new(repomocks.UsersRepo)
	ucts.producerMock = new(pubsubmock.Producer)
	ucts.ctx = context.TODO()
	config.Config.Salt = "SOME RAND SALT"

	// Define a function to assert expectations on mocks
	ucts.assertAllMocks = func() {
		ucts.usersRepositoriesMock.AssertExpectations(ucts.T())
		ucts.producerMock.AssertExpectations(ucts.T())
	}

	ucts.apiUser = &users.UsersApiImpl{}
	ucts.apiUser.SetUserMock(ucts.usersRepositoriesMock)
	ucts.apiUser.SetProducerMock(ucts.producerMock)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &UserApiTestSuite{})
}
