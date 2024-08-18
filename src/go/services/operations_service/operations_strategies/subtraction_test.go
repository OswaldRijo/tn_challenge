package operations_strategies_test

import (
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/operations_strategies"
)

func (ost *OperationStrategiesTestSuite) TestSubtractionStrategy_Apply_Success() {
	//arrange
	expectedArgs := `{"args":[1,2,3]}`
	expectedResult := "-4"
	currentUserBalance := float64(10)
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SUBTRACTION, currentUserBalance, 1, 2, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	result := ost.strategy.GetResult()
	args, errArgs := ost.strategy.GetArgsAsJson()
	cost := ost.strategy.GetCost()
	userBalanceAfterOp := ost.strategy.GetResultantUserBalance()
	_, isSubtractionStrategy := ost.strategy.(*operations_strategies.SubtractionOperationStrategy)
	ost.True(isSubtractionStrategy)
	ost.Nil(err)
	ost.Nil(errArgs)
	ost.Equal(expectedResult, result)
	ost.Equal(expectedArgs, string(args))
	ost.Equal(currentUserBalance-cost, userBalanceAfterOp)
}

func (ost *OperationStrategiesTestSuite) TestSubtractionStrategy_Apply_NotEnoughArgs() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SUBTRACTION, 10, 1)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isSubtractionStrategy := ost.strategy.(*operations_strategies.SubtractionOperationStrategy)
	ost.True(isSubtractionStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.ArgsLengthMustBeBiggerThanOne, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestSubtractionStrategy_Apply_NotEnoughBalance() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_SUBTRACTION, 0, 1, 2, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isSubtractionStrategy := ost.strategy.(*operations_strategies.SubtractionOperationStrategy)
	ost.True(isSubtractionStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.UserBalanceIsNotEnough, err.Error())
}
