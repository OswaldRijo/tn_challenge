package users_test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"truenorth/packages/utils"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/api/users"
	"truenorth/services/users_service/test"
)

func (ucts *UserApiTestSuite) Test_CheckCredentials_Success() {
	// arrange
	users.HashString = utils.HashString
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "messi", Password: "amTheBest"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(test.TheGoat, nil).Once()

	// act
	user, err := ucts.apiUser.CheckCredentials(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.NotNil(user)
	ucts.Equal(test.TheGoat.Username, user.Username)
	ucts.Equal(usersservicepb.UserStatus(test.TheGoat.Status), user.Status)
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CheckCredentials_PasswordDoesNotMatches() {
	// arrange
	users.HashString = utils.HashString
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "messi", Password: "amNotTheBest"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(test.TheGoat, nil).Once()

	// act
	user, err := ucts.apiUser.CheckCredentials(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(users.InvalidCredentials, err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CheckCredentials_UserDoesNotExist() {
	// arrange
	users.HashString = utils.HashString
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "messi", Password: "amTheBest"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, nil).Once()

	// act
	user, err := ucts.apiUser.CheckCredentials(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(users.InvalidCredentials, err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CheckCredentials_ErrorHashing() {
	// arrange
	expectedErr := errors.New("err hashing pass")
	users.HashString = func(password string, salt string) (string, error) {
		return "", expectedErr
	}
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "messi", Password: "amTheBest"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(test.TheGoat, nil).Once()

	// act
	user, err := ucts.apiUser.CheckCredentials(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CheckCredentials_ErrorGettingUser() {
	// arrange
	users.HashString = utils.HashString
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "messi", Password: "amTheBest"}
	expectedErr := errors.New("err getting user")
	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, expectedErr).Once()

	// act
	user, err := ucts.apiUser.CheckCredentials(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}
