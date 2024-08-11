package users

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"truenorth/packages/common"
	usersservicepb "truenorth/pb/users"
)

func (uc *UsersControllerImpl) GetUserByUsername(ctx context.Context, req *usersservicepb.GetUserByUsernameRequest) (*usersservicepb.GetUserByUsernameResponse, error) {
	if req.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, MissingUsername)
	}

	user, err := uc.usersApi.GetUserByUsername(ctx, req.GetUsername())

	if err != nil {
		return nil, status.Error(common.HandleApiError(err), err.Error())
	}
	return &usersservicepb.GetUserByUsernameResponse{
		User: user,
	}, nil
}
