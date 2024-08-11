package users

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	usersservicepb "truenorth/pb/users"
	"truenorth/services/users_service/models"
)

func ParseUserModelToPB(user *models.User) *usersservicepb.User {
	return &usersservicepb.User{
		Id:        int64(user.ID),
		Username:  user.Username,
		Status:    usersservicepb.UserStatus(user.Status),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
