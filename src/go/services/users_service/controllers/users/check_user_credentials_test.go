package users_test

import (
	"errors"

	"google.golang.org/grpc/status"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/controllers/users"
)

func (ucts *UserControllerTestSuite) Test_CheckUserCredentials_Success() {
	// arrange
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "some_username", Password: "some_password"}
	pbUserExpected := &usersservicepb.User{Username: req.GetUsername(), Id: mjId, Status: usersservicepb.UserStatus_ACTIVE}
	ucts.usersApiMock.On("CheckCredentials", ucts.ctx, req).Return(pbUserExpected, nil).Once()

	// act
	res, err := ucts.controller.CheckUserCredentials(ucts.ctx, req)

	// assert
	ucts.Nil(err)
	ucts.Equal(pbUserExpected, res.GetUser())
	ucts.NotNil(res.GetUser())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_CheckUserCredentials_ApiError() {
	// arrange
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "some_username", Password: "some_password"}
	errExpected := errors.New("error checking creds")
	ucts.usersApiMock.On("CheckCredentials", ucts.ctx, req).Return(nil, errExpected).Once()

	// act
	res, err := ucts.controller.CheckUserCredentials(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected.Error(), rpcStatus.Message())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_CheckUserCredentials_MissingUsername() {
	// arrange
	req := &usersservicepb.CheckUserCredentialsRequest{Password: "some_password"}
	errExpected := users.MissingUsername

	// act
	res, err := ucts.controller.CheckUserCredentials(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected, rpcStatus.Message())
	ucts.assertAllMocks()
}

func (ucts *UserControllerTestSuite) Test_CheckUserCredentials_MissingPassword() {
	// arrange
	req := &usersservicepb.CheckUserCredentialsRequest{Username: "some_username"}
	errExpected := users.MissingPassword

	// act
	res, err := ucts.controller.CheckUserCredentials(ucts.ctx, req)

	// assert
	rpcStatus, isRPCErr := status.FromError(err)
	ucts.Nil(res)
	ucts.NotNil(err)
	ucts.True(isRPCErr)
	ucts.Equal(errExpected, rpcStatus.Message())
	ucts.assertAllMocks()
}
