package users

import (
	usersservicepb "truenorth/pb/users"
	usersapi "truenorth/services/users_service/api/users"
)

type UsersControllerImpl struct {
	usersApi usersapi.UserApi
	usersservicepb.UnimplementedUserServiceServer
}

func NewUsersController() *UsersControllerImpl {
	return &UsersControllerImpl{
		usersApi: usersapi.NewUsersApi(),
	}
}
