package operations

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/models"
)

func ParseBalanceModelToPb(balance *models.Balance) *operationspb.Balance {
	return &operationspb.Balance{
		Id:             balance.ID,
		UserId:         balance.UserID,
		CurrentBalance: balance.CurrentBalance,
		CreatedAt:      timestamppb.New(balance.CreatedAt),
		UpdatedAt:      timestamppb.New(balance.UpdatedAt),
	}
}

func ParseRecordModelToPb(record *models.Record) *operationspb.Record {
	r := &operationspb.Record{
		Id:                record.ID,
		OperationId:       record.OperationID,
		UserBalance:       record.UserBalance,
		Deleted:           record.Deleted,
		OperationResponse: string(record.OperationResponse),
		CreatedAt:         timestamppb.New(record.CreatedAt),
		UpdatedAt:         timestamppb.New(record.UpdatedAt),
	}
	if record.Operation != nil {
		r.Operation = ParseOperationModelToPb(record.Operation)
	}
	return r
}

func ParseRecordModelArrToPb(records []*models.Record) []*operationspb.Record {
	recordsPbArr := make([]*operationspb.Record, len(records))
	for i, record := range records {
		recordsPbArr[i] = ParseRecordModelToPb(record)
	}
	return recordsPbArr
}

func ParseOperationModelToPb(operation *models.Operation) *operationspb.Operation {
	return &operationspb.Operation{
		Id:            operation.ID,
		UserId:        operation.UserID,
		OperationType: operation.OperationType.ToPb(),
		Cost:          operation.Cost,
		Args:          string(operation.Args),
	}
}
