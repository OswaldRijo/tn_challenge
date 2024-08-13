package users

import (
	"context"

	"gorm.io/gorm"

	"truenorth/packages/database"
	"truenorth/services/users_service/models"
)

//go:generate mockery --name=UsersRepo --output=../../../../mocks/users_service/reposirories
type UsersRepo interface {
	CreateUser(ctx context.Context, user *models.User, tx *gorm.DB) error
	GetUser(ctx context.Context, filterAttr map[string]interface{}) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type UsersRepoImpl struct {
	db *gorm.DB
}

func (u *UsersRepoImpl) UpdateUser(ctx context.Context, user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func NewUsersRepo() UsersRepo {
	return &UsersRepoImpl{
		db: database.GetInstance(),
	}
}
