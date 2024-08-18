package operations_strategies_test

import (
	operationspb "truenorth/pb/operations"
	"truenorth/services/operations_service/operations_strategies"
)

func (ost *OperationStrategiesTestSuite) TestAdditionStrategy_Apply_Success() {
	//arrange
	expectedArgs := `{"args":[1,2,3]}`
	expectedResult := "6"
	currentUserBalance := float64(10)
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_ADDITION, currentUserBalance, 1, 2, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	result := ost.strategy.GetResult()
	args, errArgs := ost.strategy.GetArgsAsJson()
	cost := ost.strategy.GetCost()
	userBalanceAfterOp := ost.strategy.GetResultantUserBalance()
	_, isAdditionStrategy := ost.strategy.(*operations_strategies.AdditionOperationStrategy)
	ost.True(isAdditionStrategy)
	ost.Nil(err)
	ost.Nil(errArgs)
	ost.Equal(expectedResult, result)
	ost.Equal(expectedArgs, string(args))
	ost.Equal(currentUserBalance-cost, userBalanceAfterOp)
}

func (ost *OperationStrategiesTestSuite) TestAdditionStrategy_Apply_NotEnoughArgs() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_ADDITION, 10, 1)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isAdditionStrategy := ost.strategy.(*operations_strategies.AdditionOperationStrategy)
	ost.True(isAdditionStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.ArgsLengthMustBeBiggerThanOne, err.Error())
}

func (ost *OperationStrategiesTestSuite) TestAdditionStrategy_Apply_NotEnoughBalance() {
	//arrange
	ost.strategy = operations_strategies.NewOperationStrategy(operationspb.OperationType_ADDITION, 0, 1, 2, 3)

	// act
	err := ost.strategy.Apply(ost.ctx)

	// assert
	_, isAdditionStrategy := ost.strategy.(*operations_strategies.AdditionOperationStrategy)
	ost.True(isAdditionStrategy)
	ost.NotNil(err)
	ost.Equal(operations_strategies.UserBalanceIsNotEnough, err.Error())
}
