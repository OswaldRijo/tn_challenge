package users

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"truenorth/packages/utils"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/config"
	"truenorth/services/users_service/models"
)

func (u *UsersApiImpl) CreateUser(ctx context.Context, userRequest *usersservicepb.CreateUserRequests) (*usersservicepb.User, error) {
	user, err := u.usersRepository.GetUser(ctx, map[string]interface{}{"username": userRequest.GetUsername()})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user != nil {
		return nil, status.Errorf(codes.InvalidArgument, UserAlreadyExistsError)
	}
	passHashed, err := utils.HashString(userRequest.GetPassword(), config.Config.Salt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	user = &models.User{
		Username:  userRequest.GetUsername(),
		Password:  passHashed,
		Status:    models.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return ParseUserModelToPB(user), nil
}
