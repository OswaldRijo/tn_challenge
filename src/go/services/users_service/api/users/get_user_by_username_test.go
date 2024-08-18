package users_test

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"truenorth/packages/utils"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/api/users"
	"truenorth/services/users_service/test"
)

func (ucts *UserApiTestSuite) Test_GetUserByUsername_Success() {
	// arrange
	req := &usersservicepb.GetUserByUsernameRequest{Username: "messi"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(test.TheGoat, nil).Once()

	// act
	user, err := ucts.apiUser.GetUserByUsername(ucts.ctx, req.GetUsername())

	// assert
	ucts.Nil(err)
	ucts.NotNil(user)
	ucts.Equal(test.TheGoat.Username, user.Username)
	ucts.Equal(usersservicepb.UserStatus(test.TheGoat.Status), user.Status)
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_GetUserByUsername_UserDoesNotExist() {
	// arrange
	req := &usersservicepb.GetUserByUsernameRequest{Username: "messi"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, nil).Once()

	// act
	user, err := ucts.apiUser.GetUserByUsername(ucts.ctx, req.GetUsername())

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(users.UserNotFoundError, err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_GetUserByUsername_ErrorGettingUser() {
	// arrange
	users.HashString = utils.HashString
	req := &usersservicepb.GetUserByUsernameRequest{Username: "messi"}
	expectedErr := errors.New("err getting user")
	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, expectedErr).Once()

	// act
	user, err := ucts.apiUser.GetUserByUsername(ucts.ctx, req.GetUsername())

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}
