package users

import (
	"context"
	"fmt"

	"truenorth/packages/common"
	"truenorth/packages/utils"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/config"
)

var HashString = utils.HashString

func (u *UsersApiImpl) CheckCredentials(ctx context.Context, userRequest *usersservicepb.CheckUserCredentialsRequest) (*usersservicepb.User, error) {
	user, err := u.usersRepository.GetUser(ctx, map[string]interface{}{"username": userRequest.GetUsername()})

	if err != nil {
		return nil, common.NewAPIErrorInternal(err)
	}

	if user == nil {
		return nil, common.NewAPIErrorInvalidArgument(fmt.Errorf(InvalidCredentials))
	}

	passHashed, err := HashString(userRequest.GetPassword(), config.Config.Salt)
	if err != nil {
		return nil, common.NewAPIErrorInternal(err)
	}

	if user.Password != passHashed {
		return nil, common.NewAPIErrorPermissionsDenied(fmt.Errorf(InvalidCredentials))
	}

	return ParseUserModelToPB(user), nil
}
