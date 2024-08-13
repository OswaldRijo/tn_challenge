package users

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"truenorth/packages/common"
	usersservicepb "truenorth/pb/users"
)

func (uc *UsersControllerImpl) CreateUser(ctx context.Context, req *usersservicepb.CreateUserRequest) (*usersservicepb.CreateUserResponse, error) {
	if req.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, MissingUsername)
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, MissingPassword)
	}

	user, err := uc.usersApi.CreateUser(ctx, req)

	if err != nil {
		return nil, status.Error(common.HandleApiError(err), err.Error())
	}
	return &usersservicepb.CreateUserResponse{
		User: user,
	}, nil
}
