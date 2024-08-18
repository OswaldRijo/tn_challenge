package operations_strategies

import (
	"context"

	"github.com/pkg/errors"

	"truenorth/services/operations_service/config"
)

type SubtractionOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (aos *SubtractionOperationStrategy) GetResult() string {
	return parseResultToString(aos.result)
}

func (aos *SubtractionOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson(aos.args...)
}

func (aos *SubtractionOperationStrategy) Apply(ctx context.Context) error {
	if len(aos.args) < 2 {
		return errors.New(ArgsLengthMustBeBiggerThanOne)
	}

	aos.result = 0
	for i, arg := range aos.args {
		if i == 0 {
			aos.result = arg
		} else {
			aos.result -= arg
		}
	}
	return aos.deductCostFromUserBalance()
}

func NewSubtractionOperationStrategy(args ...float64) *SubtractionOperationStrategy {
	return &SubtractionOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      config.Config.SubtractionOperationCost,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   args,
	}
}
