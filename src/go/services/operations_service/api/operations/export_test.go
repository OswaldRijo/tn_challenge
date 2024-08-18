package operations

import (
	"truenorth/services/operations_service/repositories/balances"
	"truenorth/services/operations_service/repositories/operations"
	"truenorth/services/operations_service/repositories/records"
)

func (uci *OperationsApiImpl) SetOperationsRepo(repo operations.OperationsRepo) {
	uci.operationsRepo = repo
}

func (uci *OperationsApiImpl) SetBalancesRepo(repo balances.BalancesRepo) {
	uci.balancesRepo = repo
}

func (uci *OperationsApiImpl) SetRecordsRepo(repo records.RecordsRepo) {
	uci.recordsRepo = repo
}
