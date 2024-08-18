package operations_strategies

import (
	"context"
	"errors"

	"truenorth/services/operations_service/config"
)

type DivisionOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (dos *DivisionOperationStrategy) GetResult() string {
	return parseResultToString(dos.result)
}

func (dos *DivisionOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson(dos.args...)
}

func (dos *DivisionOperationStrategy) Apply(ctx context.Context) error {
	if len(dos.args) < 2 {
		return errors.New(ArgsLengthMustBeBiggerThanOne)
	}

	dos.result = 0
	for i, arg := range dos.args {
		if i == 0 {
			dos.result = arg
		} else {
			if arg == 0 {
				return errors.New(ArgCantBeZero)
			}

			dos.result /= arg
		}
	}
	return dos.deductCostFromUserBalance()

}

func NewDivisionOperationStrategy(args ...float64) *DivisionOperationStrategy {
	return &DivisionOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      config.Config.DivisionOperationCost,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   args,
	}
}
