package users_test

import (
	"github.com/stretchr/testify/mock"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/models"
)

const expectedId = uint(1)

func (ucts *UserApiTestSuite) Test_CreateUser_Success() {
	// arrange
	req := &usersservicepb.CreateUserRequests{Username: "some_username", Password: "some_password"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, nil).Once()

	ucts.usersRepositoriesMock.On("CreateUser", ucts.ctx, mock.MatchedBy(func(it *models.User) bool {
		it.ID = expectedId
		return it.Username == req.GetUsername() && it.Password != ""
	})).Return(nil).Once()

	// act
	user, err := ucts.apiUser.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(expectedId, uint(user.GetId()))
	ucts.NotNil(user)
}
