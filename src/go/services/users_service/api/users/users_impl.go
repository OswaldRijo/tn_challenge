package users

import (
	"context"

	"truenorth/packages/pubsub/publisher"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/config"
	usersrepo "truenorth/services/users_service/repositories/users"
)

//go:generate mockery --name=UserApi --output=../../../../mocks/users_service/api
type UserApi interface {
	CheckCredentials(ctx context.Context, user *usersservicepb.CheckUserCredentialsRequest) (*usersservicepb.User, error)
	CreateUser(ctx context.Context, user *usersservicepb.CreateUserRequest) (*usersservicepb.User, error)
	GetUserByUsername(ctx context.Context, username string) (*usersservicepb.User, error)
	GetUser(ctx context.Context, id int64) (*usersservicepb.User, error)
	UpdateUser(ctx context.Context, user *usersservicepb.User) error
}

type UsersApiImpl struct {
	usersRepository usersrepo.UsersRepo
	producer        publisher.Producer
}

func (u *UsersApiImpl) GetUser(ctx context.Context, id int64) (*usersservicepb.User, error) {
	//TODO implement me
	return nil, nil
}

func (u *UsersApiImpl) UpdateUser(ctx context.Context, user *usersservicepb.User) error {
	//TODO implement me
	return nil
}

func NewUsersApi() UserApi {
	return &UsersApiImpl{
		usersRepository: usersrepo.NewUsersRepo(),
		producer:        publisher.NewProducer(config.Config.UserCreatedTopicArn),
	}
}
