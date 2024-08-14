package operations

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"truenorth/packages/common"
	"truenorth/packages/database"
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/models"
)

func (u *OperationsApiImpl) DeleteRecord(ctx context.Context, deleteRecordReq *operationspb.DeleteRecordsRequest) ([]*operationspb.Record, *operationspb.Balance, error) {
	records, err := u.recordsRepo.FindRecordsByIds(ctx, deleteRecordReq.GetRecordIds()...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, common.NewAPIErrorResourceNotFound(err)
		}
		return nil, nil, common.NewAPIErrorInternal(err)

	}

	userBalance, err := u.balancesRepo.GetBalanceByUserId(ctx, deleteRecordReq.GetUserId())
	if err != nil {
		return nil, nil, common.NewAPIErrorInternal(err)
	}

	err = u.checkRecordsBelongsToUser(records, userBalance)
	if err != nil {
		return nil, nil, err
	}

	userBalance.CurrentBalance += u.getTotalBalanceToRefund(records)

	err = database.PerformDbTransaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		for _, r := range records {
			err = u.recordsRepo.DeleteRecordById(ctx, r.ID, tx)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return common.NewAPIErrorResourceNotFound(err)
				}
				return common.NewAPIErrorInternal(err)
			}
		}

		err = u.balancesRepo.UpdateBalance(ctx, userBalance, tx)
		if err != nil {
			return common.NewAPIErrorInternal(err)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return ParseRecordModelArrToPb(records), ParseBalanceModelToPb(userBalance), nil
}

func (u *OperationsApiImpl) checkRecordsBelongsToUser(records []*models.Record, balance *models.Balance) error {
	for _, record := range records {
		if record.UserID != balance.UserID {
			return common.NewAPIErrorPermissionsDenied(fmt.Errorf(RecordDoesNotBelongToTheUser))
		}
	}
	return nil
}

func (u *OperationsApiImpl) getTotalBalanceToRefund(records []*models.Record) float64 {
	totalBalance := 0.0
	for _, record := range records {
		totalBalance += record.Operation.Cost
	}

	return totalBalance
}
