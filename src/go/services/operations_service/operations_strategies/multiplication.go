package operations_strategies

type MultiplicationOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (mos *MultiplicationOperationStrategy) setArgs(args ...float64) *MultiplicationOperationStrategy {
	mos.args = args
	return mos
}

func (mos *MultiplicationOperationStrategy) GetResultAsJson() ([]byte, error) {
	return serializeResultAsJson(mos.result)
}

func (mos *MultiplicationOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeResultAsJson(mos.args)
}

func (mos *MultiplicationOperationStrategy) Apply() error {
	mos.result = 0
	for _, arg := range mos.args {
		mos.result *= arg
	}
	mos.deductCostFromUserBalance()

	return nil
}

func NewMultiplicationOperationStrategy() *MultiplicationOperationStrategy {
	return &MultiplicationOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      100,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   []float64{},
	}
}
