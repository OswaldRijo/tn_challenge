package users

import (
	usersservicepb "truenorth/pb/users"
)

type UsersControllerImpl struct {
	usersservicepb.UnimplementedUserServiceServer
}

func NewUsersController() *UsersControllerImpl {
	return &UsersControllerImpl{}
}
