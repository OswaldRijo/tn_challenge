package operations

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"truenorth/packages/database"
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/models"
	operationsstrategies "truenorth/services/operations_service/operations_strategies"
)

func (u *OperationsApiImpl) ApplyOperation(ctx context.Context, operationReq *operationspb.ApplyOperationRequest) (*operationspb.Operation, *operationspb.Record, *operationspb.Balance, error) {
	userBalance, err := u.balancesRepo.GetBalanceByUserId(ctx, operationReq.GetUserId())
	if err != nil {
		return nil, nil, nil, status.Error(codes.Internal, err.Error())
	}

	operationStrategy := operationsstrategies.NewOperationStrategy(operationReq.GetOperationType(), userBalance.CurrentBalance, operationReq.Args...)
	err = operationStrategy.Apply()
	if err != nil {
		return nil, nil, nil, status.Error(codes.Internal, err.Error())
	}

	args, err := operationStrategy.GetArgsAsJson()
	if err != nil {
		return nil, nil, nil, status.Error(codes.Internal, err.Error())
	}

	result, err := operationStrategy.GetResultAsJson()

	if err != nil {
		return nil, nil, nil, status.Error(codes.Internal, err.Error())
	}

	operationCost := operationStrategy.GetCost()
	currentUserBalance := operationStrategy.GetResultantUserBalance()
	var operationModel *models.Operation
	var recordModel *models.Record
	err = database.PerformDbTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		now := time.Now()
		operationModel, err = u.insertOperationModel(ctx, operationReq, now, operationCost, args, tx)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		recordModel, err = u.insertRecordModel(ctx, operationReq, operationModel, now, currentUserBalance, result, tx)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		userBalance.CurrentBalance = currentUserBalance
		err = u.balancesRepo.UpdateBalance(ctx, userBalance, tx)

		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		return nil
	})

	if err != nil {
		return nil, nil, nil, err
	}

	return ParseOperationModelToPb(operationModel), ParseRecordModelToPb(recordModel), ParseBalanceModelToPb(userBalance), nil
}

func (u *OperationsApiImpl) insertOperationModel(ctx context.Context, operationReq *operationspb.ApplyOperationRequest, now time.Time, cost float64, args []byte, tx *gorm.DB) (*models.Operation, error) {
	operationModel := models.NewOperation().SetCreatedAt(now).
		SetUpdatedAt(now).
		SetUserID(operationReq.GetUserId()).
		SetCost(cost).
		SetOperationType(models.OpTypeFromPb(operationReq.GetOperationType())).
		SetArgs(args)

	err := u.operationsRepo.CreateOperation(ctx, operationModel, tx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return operationModel, nil
}

func (u *OperationsApiImpl) insertRecordModel(ctx context.Context, operationReq *operationspb.ApplyOperationRequest, operationModel *models.Operation, now time.Time, userBalance float64, result []byte, tx *gorm.DB) (*models.Record, error) {
	recordModel := models.NewRecord().SetCreatedAt(now).
		SetUpdatedAt(now).
		SetOperationID(operationModel.ID).
		SetUserID(operationReq.GetUserId()).
		SetUserBalance(userBalance).
		SetDeleted(false).
		SetOperationResponse(result)

	err := u.recordsRepo.CreateRecord(ctx, recordModel, tx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return recordModel, nil
}
