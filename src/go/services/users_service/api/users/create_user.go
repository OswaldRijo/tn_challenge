package users

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"truenorth/packages/common"
	"truenorth/packages/database"
	"truenorth/packages/pubsub/publisher"
	"truenorth/packages/utils"
	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/config"
	"truenorth/services/users_service/models"
)

func (u *UsersApiImpl) CreateUser(ctx context.Context, userRequest *usersservicepb.CreateUserRequest) (*usersservicepb.User, error) {
	user, err := u.usersRepository.GetUser(ctx, map[string]interface{}{"username": userRequest.GetUsername()})

	if err != nil {
		return nil, common.NewAPIErrorInternal(err)
	}

	if user != nil {
		return nil, common.NewAPIErrorInvalidArgument(fmt.Errorf(UserAlreadyExistsError))
	}
	passHashed, err := utils.HashString(userRequest.GetPassword(), config.Config.Salt)
	if err != nil {
		return nil, common.NewAPIErrorInternal(err)
	}
	user = &models.User{
		Username:  userRequest.GetUsername(),
		Password:  passHashed,
		Status:    models.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = database.PerformDbTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		err = u.usersRepository.CreateUser(ctx, user, tx)
		if err != nil {
			return common.NewAPIErrorInternal(err)
		}

		messageMap := make(map[string]interface{})
		messageMap["user_id"] = fmt.Sprintf("%v", user.ID)

		producer := publisher.NewProducer(config.Config.UserCreatedTopicArn)

		err = producer.SendMessage(ctx, messageMap)
		if err != nil {
			return common.NewAPIErrorInternal(err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ParseUserModelToPB(user), nil
}
