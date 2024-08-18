package operations_strategies_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"truenorth/services/operations_service/config"
	"truenorth/services/operations_service/operations_strategies"
)

type OperationStrategiesTestSuite struct {
	suite.Suite

	strategy operations_strategies.OperationStrategy
	ctx      context.Context
}

func (ost *OperationStrategiesTestSuite) SetupTest() {
	// Initialize necessary dependencies
	ost.ctx = context.TODO()
	config.Config.AdditionOperationCost = 5
	config.Config.SubtractionOperationCost = 6
	config.Config.DivisionOperationCost = 7
	config.Config.MultiplicationOperationCost = 8
	config.Config.RandomStringOperationCost = 9
	config.Config.SquareRootOperationCost = 10
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, &OperationStrategiesTestSuite{})
}
