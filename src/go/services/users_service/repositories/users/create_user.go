package users

import (
	"context"

	"gorm.io/gorm"

	"truenorth/services/users_service/models"
)

func (u *UsersRepoImpl) CreateUser(ctx context.Context, user *models.User, tx *gorm.DB) error {
	result := tx.WithContext(ctx).Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
