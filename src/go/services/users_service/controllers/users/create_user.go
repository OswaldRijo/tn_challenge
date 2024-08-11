package users

import (
	"golang.org/x/net/context"
	usersservicepb "truenorth/pb/users"
)

func (*UsersControllerImpl) CreateUser(ctx context.Context, req *usersservicepb.CreateUserRequests) *usersservicepb.CreateUserResponse {

	return &usersservicepb.CreateUserResponse{}
}
