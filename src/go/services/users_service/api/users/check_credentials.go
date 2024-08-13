package users

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"truenorth/packages/utils"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/config"
)

func (u *UsersApiImpl) CheckCredentials(ctx context.Context, userRequest *usersservicepb.CheckUserCredentialsRequest) (*usersservicepb.User, error) {
	user, err := u.usersRepository.GetUser(ctx, map[string]interface{}{"username": userRequest.GetUsername()})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, InvalidCredentials)
	}

	passHashed, err := utils.HashString(userRequest.GetPassword(), config.Config.Salt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user.Password != passHashed {
		return nil, status.Errorf(codes.PermissionDenied, InvalidCredentials)
	}

	return ParseUserModelToPB(user), nil
}
