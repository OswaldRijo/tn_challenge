package operations_strategies_test

import (
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/operations_strategies"
)

func (ost *OperationStrategiesTestSuite) TestSquareRootStrategy_Apply_Success() {
	//arrange
	expectedArgs := `{"args":[4]}`
	expectedResult := "2"
	currentUserBalance := float64(10)
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SQUARE_ROOT, currentUserBalance, 4)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	result := ost.strategy.GetResult()
	args, errArgs := ost.strategy.GetArgsAsJson()
	cost := ost.strategy.GetCost()
	userBalanceAfterOp := ost.strategy.GetResultantUserBalance()
	_, isSquareRootStrategy := ost.strategy.(*operations_strategies.SquareRootOperationStrategy)
	ost.True(isSquareRootStrategy)
	ost.Nil(err)
	ost.Nil(errArgs)
	ost.Equal(expectedResult, result)
	ost.Equal(expectedArgs, string(args))
	ost.Equal(currentUserBalance-cost, userBalanceAfterOp)
}

func (ost *OperationStrategiesTestSuite) TestSquareRootStrategy_Apply_ArgsQtyMustBeOne() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SQUARE_ROOT, 10, 1, 1)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isSquareRootStrategy := ost.strategy.(*operations_strategies.SquareRootOperationStrategy)
	ost.True(isSquareRootStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.ArgsLengthMustBeOne, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestSquareRootStrategy_Apply_ArgsIsNegative() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SQUARE_ROOT, 10, -0.5)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isSquareRootStrategy := ost.strategy.(*operations_strategies.SquareRootOperationStrategy)
	ost.True(isSquareRootStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.NegativeNumbersNotAllowed, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestSquareRootStrategy_Apply_NotEnoughBalance() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SQUARE_ROOT, 0, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	ost.NotNil(err)
	ost.Equal(operations_strategies.UserBalanceIsNotEnough, err.Error())
}
