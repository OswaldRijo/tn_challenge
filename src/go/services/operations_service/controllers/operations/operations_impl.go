package operations

import (
	"context"

	operationspb "truenorth/pb/operations"
)

type OperationsControllerImpl struct {
	operationspb.UnimplementedOperationsServiceServer
}

func (s *OperationsControllerImpl) ApplyOperation(ctx context.Context, request *operationspb.ApplyOperationRequest) (*operationspb.ApplyOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OperationsControllerImpl) GetUserBalance(ctx context.Context, request *operationspb.GetUserBalanceRequest) (*operationspb.GetUserBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OperationsControllerImpl) FilterRecords(ctx context.Context, request *operationspb.FilterRecordsRequest) (*operationspb.FilterRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OperationsControllerImpl) DeleteRecords(ctx context.Context, request *operationspb.DeleteRecordsRequest) (*operationspb.DeleteRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewOperationsController() *OperationsControllerImpl {
	return &OperationsControllerImpl{}
}
