package operations

import (
	"truenorth/services/operations_service/api/operations"
)

func (uci *OperationsControllerImpl) SetOperationsApi(api operations.OperationsApi) {
	uci.operationsApi = api
}
