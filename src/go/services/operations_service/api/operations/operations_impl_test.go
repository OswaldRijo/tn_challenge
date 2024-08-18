package operations_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	repomocks "truenorth/mocks/operations_service/repositories"
	"truenorth/services/operations_service/api/operations"
)

type OperationApiTestSuite struct {
	suite.Suite

	api                *operations.OperationsApiImpl
	ctx                context.Context
	operationsRepoMock *repomocks.OperationsRepo
	balancesRepoMock   *repomocks.BalancesRepo
	recordsRepoMock    *repomocks.RecordsRepo
	assertAllMocks     func()
}

func (ucts *OperationApiTestSuite) SetupTest() {
	// Initialize necessary dependencies
	ucts.operationsRepoMock = new(repomocks.OperationsRepo)
	ucts.balancesRepoMock = new(repomocks.BalancesRepo)
	ucts.recordsRepoMock = new(repomocks.RecordsRepo)
	ucts.ctx = context.TODO()

	// Define a function to assert expectations on mocks
	ucts.assertAllMocks = func() {
		ucts.operationsRepoMock.AssertExpectations(ucts.T())
		ucts.balancesRepoMock.AssertExpectations(ucts.T())
		ucts.recordsRepoMock.AssertExpectations(ucts.T())
	}

	ucts.api = &operations.OperationsApiImpl{}
	ucts.api.SetOperationsRepo(ucts.operationsRepoMock)
	ucts.api.SetBalancesRepo(ucts.balancesRepoMock)
	ucts.api.SetRecordsRepo(ucts.recordsRepoMock)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &OperationApiTestSuite{})
}
