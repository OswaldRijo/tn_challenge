package operations

import (
	"context"

	operationspb "truenorth/pb/operations"
	balancesrepo "truenorth/services/operations_service/repositories/balances"
	operationsrepo "truenorth/services/operations_service/repositories/operations"
	recordsrepo "truenorth/services/operations_service/repositories/records"
)

//go:generate mockery --name=OperationsApi --output=../../../../mocks/operations_service/api
type OperationsApi interface {
	FilterRecords(ctx context.Context, operationReq *operationspb.FilterRecordsRequest) ([]*operationspb.Record, int64, error)
	ApplyOperation(ctx context.Context, operationReq *operationspb.ApplyOperationRequest) (*operationspb.Operation, *operationspb.Record, *operationspb.Balance, error)
	DeleteRecord(ctx context.Context, deleteRecordReq *operationspb.DeleteRecordsRequest) ([]*operationspb.Record, *operationspb.Balance, error)
	CreateUserBalance(ctx context.Context, userId int64) (*operationspb.Balance, error)
	GetUserBalance(ctx context.Context, userId int64) (*operationspb.Balance, error)
}

type OperationsApiImpl struct {
	operationsRepo operationsrepo.OperationsRepo
	balancesRepo   balancesrepo.BalancesRepo
	recordsRepo    recordsrepo.RecordsRepo
}

func NewOperationsApi() OperationsApi {
	return &OperationsApiImpl{
		operationsRepo: operationsrepo.NewOperationsRepo(),
		balancesRepo:   balancesrepo.NewBalancesRepo(),
		recordsRepo:    recordsrepo.NewRecordsRepo(),
	}
}
