package users_test

import (
	"errors"

	"google.golang.org/grpc/status"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/controllers/users"
)

func (ucts *UserControllerTestSuite) Test_GetUserByUsername_Success() {
	// arrange
	req := &usersservicepb.GetUserByUsernameRequest{Username: "some_username"}
	pbUserExpected := &usersservicepb.User{Username: req.GetUsername(), Id: mjId, Status: usersservicepb.UserStatus_ACTIVE}
	ucts.usersApiMock.On("GetUserByUsername", ucts.ctx, req.GetUsername()).Return(pbUserExpected, nil).Once()

	// act
	res, err := ucts.controller.GetUserByUsername(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(pbUserExpected, res.GetUser())
	ucts.NotNil(res.GetUser())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_GetUserByUsername_ApiError() {
	// arrange
	req := &usersservicepb.GetUserByUsernameRequest{Username: "some_username"}
	errExpected := errors.New("error getting user")
	ucts.usersApiMock.On("GetUserByUsername", ucts.ctx, req.GetUsername()).Return(nil, errExpected).Once()

	// act
	res, err := ucts.controller.GetUserByUsername(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected.Error(), rpcStatus.Message())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_GetUserByUsername_MissingUsername() {
	// arrange
	req := &usersservicepb.GetUserByUsernameRequest{}
	errExpected := users.MissingUsername

	// act
	res, err := ucts.controller.GetUserByUsername(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected, rpcStatus.Message())
	ucts.assertAllMocks()
}
