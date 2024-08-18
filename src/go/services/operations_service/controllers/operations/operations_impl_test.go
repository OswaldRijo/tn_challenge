package operations_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	apimocks "truenorth/mocks/operations_service/api"
	"truenorth/services/operations_service/controllers/operations"
)

type OperationControllerTestSuite struct {
	suite.Suite

	controller        *operations.OperationsControllerImpl
	ctx               context.Context
	operationsApiMock *apimocks.OperationsApi
	assertAllMocks    func()
}

func (ucts *OperationControllerTestSuite) SetupTest() {

	// Initialize necessary dependencies
	ucts.operationsApiMock = new(apimocks.OperationsApi)
	ucts.ctx = context.TODO()

	// Define a function to assert expectations on mocks
	ucts.assertAllMocks = func() {
		ucts.operationsApiMock.AssertExpectations(ucts.T())
	}

	ucts.controller = &operations.OperationsControllerImpl{}
	ucts.controller.SetOperationsApi(ucts.operationsApiMock)
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &OperationControllerTestSuite{})
}
