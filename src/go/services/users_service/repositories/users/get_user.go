package users

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"truenorth/services/users_service/models"
)

func (u *UsersRepoImpl) GetUser(ctx context.Context, filterAttributes map[string]interface{}) (*models.User, error) {
	user := models.User{}
	result := u.db.Where(filterAttributes).First(&user).WithContext(ctx)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, nil
}
