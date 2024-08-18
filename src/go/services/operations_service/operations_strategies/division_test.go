package operations_strategies_test

import (
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/operations_strategies"
)

func (ost *OperationStrategiesTestSuite) TestDivisionStrategy_Apply_Success() {
	//arrange
	expectedArgs := `{"args":[36,2]}`
	expectedResult := "18"
	currentUserBalance := float64(10)
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_DIVISION, currentUserBalance, 36, 2)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	result := ost.strategy.GetResult()
	args, errArgs := ost.strategy.GetArgsAsJson()
	cost := ost.strategy.GetCost()
	userBalanceAfterOp := ost.strategy.GetResultantUserBalance()
	_, isDivisionStrategy := ost.strategy.(*operations_strategies.DivisionOperationStrategy)
	ost.True(isDivisionStrategy)
	ost.Nil(err)
	ost.Nil(errArgs)
	ost.Equal(expectedResult, result)
	ost.Equal(expectedArgs, string(args))
	ost.Equal(currentUserBalance-cost, userBalanceAfterOp)
}

func (ost *OperationStrategiesTestSuite) TestDivisionStrategy_Apply_NotEnoughArgs() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_DIVISION, 10, 1)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isDivisionStrategy := ost.strategy.(*operations_strategies.DivisionOperationStrategy)
	ost.True(isDivisionStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.ArgsLengthMustBeBiggerThanOne, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestDivisionStrategy_Apply_InvalidArg() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_DIVISION, 10, 5, 0)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isDivisionStrategy := ost.strategy.(*operations_strategies.DivisionOperationStrategy)
	ost.True(isDivisionStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.ArgCantBeZero, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestDivisionStrategy_Apply_NotEnoughBalance() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_DIVISION, 0, 1, 2, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	ost.NotNil(err)
	ost.Equal(operations_strategies.UserBalanceIsNotEnough, err.Error())
}
