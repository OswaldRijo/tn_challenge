package test

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	operationspb "truenorth/pb/operations"
	operationsmodels "truenorth/services/operations_service/models"
	usermodels "truenorth/services/users_service/models"
)

var RandSalt = "SOME RAND SALT"

var TheGoat = &usermodels.User{
	Username:  "messi",
	Password:  "YeVYjDfUo8B9U7j990C1ar8zONs7i6a4L8sq11cuSb8=",
	Status:    usermodels.StatusActive,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var YagamiLightBalanceModel = &operationsmodels.Balance{
	ID:             16,
	UserID:         16,
	CurrentBalance: 999999,
	CreatedAt:      time.Now(),
	UpdatedAt:      time.Now(),
}

var YagamiLightBalancePb = &operationspb.Balance{
	Id:             YagamiLightBalanceModel.ID,
	UserId:         YagamiLightBalanceModel.UserID,
	CurrentBalance: YagamiLightBalanceModel.CurrentBalance,
	CreatedAt:      timestamppb.New(time.Now()),
	UpdatedAt:      timestamppb.New(time.Now()),
}

var OperationModel = &operationsmodels.Operation{
	ID:            1,
	UserID:        YagamiLightBalanceModel.UserID,
	OperationType: operationsmodels.OperationTypeAddition,
	Cost:          0,
	CreatedAt:     time.Now(),
	UpdatedAt:     time.Now(),
}

var OperationPb = &operationspb.Operation{
	Id:            1,
	UserId:        1,
	OperationType: operationspb.OperationType_ADDITION,
	Cost:          0,
	Args:          `{"args":[1,2]}`,
	CreatedAt:     timestamppb.New(time.Now()),
	UpdatedAt:     timestamppb.New(time.Now()),
}

var RecordPb = &operationspb.Record{
	Id:                1,
	OperationId:       1,
	UserBalance:       999999,
	Deleted:           false,
	OperationResponse: "3",
	CreatedAt:         timestamppb.New(time.Now()),
	UpdatedAt:         timestamppb.New(time.Now()),
	Operation:         OperationPb,
}

var RecordsPb = []*operationspb.Record{RecordPb}

var RecordModel = &operationsmodels.Record{
	ID:                1,
	OperationID:       1,
	UserID:            YagamiLightBalanceModel.UserID,
	UserBalance:       999999,
	Deleted:           false,
	OperationResponse: "3",
	CreatedAt:         time.Now(),
	UpdatedAt:         time.Now(),
	Operation:         *OperationModel,
}

var RecordsModel = []*operationsmodels.Record{RecordModel}
