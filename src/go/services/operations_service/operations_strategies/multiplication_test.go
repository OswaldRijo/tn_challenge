package operations_strategies_test

import (
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/operations_strategies"
)

func (ost *OperationStrategiesTestSuite) TestMultiplicationStrategy_Apply_Success() {
	//arrange
	expectedArgs := `{"args":[5.5,2]}`
	expectedResult := "11"
	currentUserBalance := float64(10)
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_MULTIPLICATION, currentUserBalance, 5.5, 2)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	result := ost.strategy.GetResult()
	args, errArgs := ost.strategy.GetArgsAsJson()
	cost := ost.strategy.GetCost()
	userBalanceAfterOp := ost.strategy.GetResultantUserBalance()
	_, isMultiplicationStrategy := ost.strategy.(*operations_strategies.MultiplicationOperationStrategy)
	ost.True(isMultiplicationStrategy)
	ost.Nil(err)
	ost.Nil(errArgs)
	ost.Equal(expectedResult, result)
	ost.Equal(expectedArgs, string(args))
	ost.Equal(currentUserBalance-cost, userBalanceAfterOp)
}

func (ost *OperationStrategiesTestSuite) TestMultiplicationStrategy_Apply_NotEnoughArgs() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_MULTIPLICATION, 10, 1)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isMultiplicationStrategy := ost.strategy.(*operations_strategies.MultiplicationOperationStrategy)
	ost.True(isMultiplicationStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.ArgsLengthMustBeBiggerThanOne, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestMultiplicationStrategy_Apply_NotEnoughBalance() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_MULTIPLICATION, 0, 1, 2, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	ost.NotNil(err)
	ost.Equal(operations_strategies.UserBalanceIsNotEnough, err.Error())
}
