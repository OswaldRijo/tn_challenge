package users

import (
	"context"
	"fmt"

	"truenorth/packages/common"
	usersservicepb "truenorth/pb/users"
)

func (u *UsersApiImpl) GetUserByUsername(ctx context.Context, username string) (*usersservicepb.User, error) {
	user, err := u.usersRepository.GetUser(ctx, map[string]interface{}{"username": username})

	if err != nil {
		return nil, common.NewAPIErrorInternal(err)
	}

	if user == nil {
		return nil, common.NewAPIErrorInvalidArgument(fmt.Errorf(UserNotFoundError))
	}

	return ParseUserModelToPB(user), nil
}
