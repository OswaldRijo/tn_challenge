package users_test

import (
	"context"
	"errors"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/api/users"
	"truenorth/services/users_service/models"
	"truenorth/services/users_service/test"
)

const expectedId = int64(1)

func (ucts *UserApiTestSuite) Test_CreateUser_Success() {
	// arrange
	users.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}

	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, nil).Once()

	ucts.usersRepositoriesMock.On("CreateUser", ucts.ctx, mock.MatchedBy(func(it *models.User) bool {
		it.ID = expectedId
		return it.Username == req.GetUsername() && it.Password != ""
	}), mock.Anything).Return(nil).Once()

	ucts.producerMock.On("SendMessage", ucts.ctx, mock.Anything).Return(nil).Once()

	// act
	user, err := ucts.apiUser.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(expectedId, int64(user.GetId()))
	ucts.NotNil(user)
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CreateUser_ErrorGettingUser() {
	// arrange
	users.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}
	expectedErr := errors.New("something went wrong")
	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, expectedErr).Once()

	// act
	user, err := ucts.apiUser.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CreateUser_ErrorUserExists() {
	// arrange
	users.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}
	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(test.TheGoat, nil).Once()

	// act
	user, err := ucts.apiUser.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(users.UserAlreadyExistsError, err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CreateUser_ErrorCreatingUser() {
	// arrange
	users.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}
	expectedErr := errors.New("something went wrong")
	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, nil).Once()

	ucts.usersRepositoriesMock.On("CreateUser", ucts.ctx, mock.MatchedBy(func(it *models.User) bool {
		it.ID = expectedId
		return it.Username == req.GetUsername() && it.Password != ""
	}), mock.Anything).Return(expectedErr).Once()

	// act
	user, err := ucts.apiUser.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}

func (ucts *UserApiTestSuite) Test_CreateUser_ErrorNotifyingUserCreation() {
	// arrange
	users.InitTransaction = func(ctx context.Context, onRunningTxn func(ctx context.Context, tx *gorm.DB) error) error {
		return onRunningTxn(ctx, nil)
	}

	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}
	expectedErr := errors.New("something went wrong")
	ucts.usersRepositoriesMock.On("GetUser", ucts.ctx, mock.MatchedBy(func(it map[string]interface{}) bool {
		return it["username"] == req.GetUsername()
	})).Return(nil, nil).Once()

	ucts.usersRepositoriesMock.On("CreateUser", ucts.ctx, mock.MatchedBy(func(it *models.User) bool {
		it.ID = expectedId
		return it.Username == req.GetUsername() && it.Password != ""
	}), mock.Anything).Return(nil).Once()

	ucts.producerMock.On("SendMessage", ucts.ctx, mock.Anything).Return(expectedErr).Once()

	// act
	user, err := ucts.apiUser.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(user)
	ucts.NotNil(err)
	ucts.Equal(expectedErr.Error(), err.Error())
	ucts.assertAllMocks()
}
