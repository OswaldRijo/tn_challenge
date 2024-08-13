package operations

import (
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/api/operations"
)

type OperationsControllerImpl struct {
	operationspb.UnimplementedOperationsServiceServer
	operationsApi operations.OperationsApi
}

func NewOperationsController() *OperationsControllerImpl {
	return &OperationsControllerImpl{
		operationsApi: operations.NewOperationsApi(),
	}
}
