package users

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersservicepb "truenorth/pb/users"
)

func (u *UsersApiImpl) GetUserByUsername(ctx context.Context, username string) (*usersservicepb.User, error) {
	user, err := u.usersRepository.GetUser(ctx, map[string]interface{}{"username": username})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, UserNotFoundError)
	}

	return ParseUserModelToPB(user), nil
}
