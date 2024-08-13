package operations_strategies

type AdditionOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (aos *AdditionOperationStrategy) setArgs(args ...float64) *AdditionOperationStrategy {
	aos.args = args
	return aos
}

func (aos *AdditionOperationStrategy) GetResultAsJson() ([]byte, error) {
	return serializeResultAsJson(aos.result)
}

func (aos *AdditionOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeResultAsJson(aos.args)
}

func (aos *AdditionOperationStrategy) Apply() error {
	aos.result = 0
	for _, arg := range aos.args {
		aos.result += arg
	}
	aos.deductCostFromUserBalance()
	return nil
}

func NewAdditionOperationStrategy() *AdditionOperationStrategy {
	return &AdditionOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      100,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   []float64{},
	}
}
