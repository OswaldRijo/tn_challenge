package operations_strategies

import "math"

type SquareRootOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   float64
}

func (sros *SquareRootOperationStrategy) setArgs(args float64) *SquareRootOperationStrategy {
	sros.args = args
	return sros
}

func (sros *SquareRootOperationStrategy) GetResultAsJson() ([]byte, error) {
	return serializeResultAsJson(sros.result)
}

func (sros *SquareRootOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeResultAsJson(sros.args)
}

func (sros *SquareRootOperationStrategy) Apply() error {
	sros.result = math.Sqrt(sros.args)
	sros.deductCostFromUserBalance()
	return nil
}

func NewSquareRootOperationStrategy() *SquareRootOperationStrategy {
	return &SquareRootOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      100,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   0,
	}
}
