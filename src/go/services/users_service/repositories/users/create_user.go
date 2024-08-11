package users

import (
	"context"

	"truenorth/services/users_service/models"
)

func (u *UsersRepoImpl) CreateUser(ctx context.Context, user *models.User) error {
	result := u.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
