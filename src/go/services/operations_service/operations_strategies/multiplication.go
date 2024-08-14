package operations_strategies

import (
	"context"
	"errors"

	"truenorth/services/operations_service/config"
)

type MultiplicationOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (mos *MultiplicationOperationStrategy) GetResult() string {
	return parseResultToString(mos.result)
}

func (mos *MultiplicationOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson(mos.args...)
}

func (mos *MultiplicationOperationStrategy) Apply(ctx context.Context) error {
	if len(mos.args) < 2 {
		return errors.New(ArgsLengthMustBeBiggerThanOne)
	}

	mos.result = 1
	for _, arg := range mos.args {
		mos.result *= arg
	}
	mos.deductCostFromUserBalance()

	return nil
}

func NewMultiplicationOperationStrategy(args ...float64) *MultiplicationOperationStrategy {
	return &MultiplicationOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      config.Config.MultiplicationOperationCost,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   args,
	}
}
