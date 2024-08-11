package users

import (
	"context"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/models"
	usersrepo "truenorth/services/users_service/repositories/users"
)

//go:generate mockery --name=UserApi --output=../../../../mocks/users_service/api
type UserApi interface {
	CreateUser(ctx context.Context, user *usersservicepb.CreateUserRequests) (*usersservicepb.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type UsersApiImpl struct {
	usersRepository usersrepo.UsersRepo
}

func (u *UsersApiImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	//TODO implement me
	return nil, nil
}

func (u *UsersApiImpl) GetUser(ctx context.Context, id int64) (*models.User, error) {
	//TODO implement me
	return nil, nil
}

func (u *UsersApiImpl) UpdateUser(ctx context.Context, user *models.User) error {
	//TODO implement me
	return nil
}

func NewUsersApi() UserApi {
	return &UsersApiImpl{
		usersRepository: usersrepo.NewUsersRepo(),
	}
}
