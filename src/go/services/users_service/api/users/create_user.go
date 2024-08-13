package users

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

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
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if user != nil {
		return nil, status.Errorf(codes.InvalidArgument, UserAlreadyExistsError)
	}
	passHashed, err := utils.HashString(userRequest.GetPassword(), config.Config.Salt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
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
			return status.Errorf(codes.Internal, err.Error())
		}

		messageMap := make(map[string]interface{})
		messageMap["user_id"] = user.ID

		producer := publisher.NewProducer(config.Config.UserCreatedTopicArn)

		err = producer.SendMessage(ctx, messageMap)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ParseUserModelToPB(user), nil
}
