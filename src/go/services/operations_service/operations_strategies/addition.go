package operations_strategies

import (
	"context"
	"errors"

	"truenorth/services/operations_service/config"
)

type AdditionOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (aos *AdditionOperationStrategy) GetResult() string {
	return parseResultToString(aos.result)
}

func (aos *AdditionOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson(aos.args...)
}

func (aos *AdditionOperationStrategy) Apply(ctx context.Context) error {
	if len(aos.args) < 2 {
		return errors.New(ArgsLengthMustBeBiggerThanOne)
	}

	aos.result = 0
	for _, arg := range aos.args {
		aos.result += arg
	}
	return aos.deductCostFromUserBalance()
}

func NewAdditionOperationStrategy(args ...float64) *AdditionOperationStrategy {
	return &AdditionOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      config.Config.AdditionOperationCost,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   args,
	}
}
