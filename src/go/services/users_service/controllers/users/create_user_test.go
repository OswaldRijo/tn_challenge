package users_test

import (
	"errors"

	"google.golang.org/grpc/status"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/controllers/users"
)

const mjId = 23

func (ucts *UserControllerTestSuite) Test_CreateUser_Success() {
	// arrange
	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}
	pbUserExpected := &usersservicepb.User{Username: req.GetUsername(), Id: mjId, Status: usersservicepb.UserStatus_ACTIVE}
	ucts.usersApiMock.On("CreateUser", ucts.ctx, req).Return(pbUserExpected, nil).Once()

	// act
	res, err := ucts.controller.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(pbUserExpected, res.GetUser())
	ucts.NotNil(res.GetUser())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_CreateUser_ApiErrorCreatingUser() {
	// arrange
	req := &usersservicepb.CreateUserRequest{Username: "some_username", Password: "some_password"}
	expectedErr := errors.New("error creating user")
	ucts.usersApiMock.On("CreateUser", ucts.ctx, req).Return(nil, expectedErr).Once()

	// act
	res, err := ucts.controller.CreateUser(ucts.ctx, req)

	// assert
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.Contains(err.Error(), expectedErr.Error())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_CreateUser_MissingUsername() {
	// arrange
	req := &usersservicepb.CreateUserRequest{Password: "some_password"}
	errExpected := users.MissingUsername

	// act
	res, err := ucts.controller.CreateUser(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected, rpcStatus.Message())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_CreateUser_MissingPassword() {
	// arrange
	req := &usersservicepb.CreateUserRequest{Username: "some_username"}
	errExpected := users.MissingPassword

	// act
	res, err := ucts.controller.CreateUser(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected, rpcStatus.Message())
	ucts.assertAllMocks()
}
