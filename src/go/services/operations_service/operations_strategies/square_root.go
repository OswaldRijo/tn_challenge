package operations_strategies

import (
	"context"
	"errors"
	"math"

	"truenorth/services/operations_service/config"
)

type SquareRootOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (sros *SquareRootOperationStrategy) GetResult() string {
	return parseResultToString(sros.result)
}

func (sros *SquareRootOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson(sros.args...)
}

func (sros *SquareRootOperationStrategy) Apply(ctx context.Context) error {
	if len(sros.args) != 1 {
		return errors.New(ArgsLengthMustBeOne)
	}

	if sros.args[0] < 0 {
		return errors.New(NegativeNumbersNotAllowed)
	}
	sros.result = math.Sqrt(sros.args[0])
	return sros.deductCostFromUserBalance()
}

func NewSquareRootOperationStrategy(arg ...float64) *SquareRootOperationStrategy {
	return &SquareRootOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      config.Config.SquareRootOperationCost,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   arg,
	}
}
