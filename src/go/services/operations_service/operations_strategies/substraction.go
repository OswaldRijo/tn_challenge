package operations_strategies

type SubtractionOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (aos *SubtractionOperationStrategy) setArgs(args ...float64) *SubtractionOperationStrategy {
	aos.args = args
	return aos
}

func (aos *SubtractionOperationStrategy) GetResultAsJson() ([]byte, error) {
	return serializeResultAsJson(aos.result)
}

func (aos *SubtractionOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson(aos.args)
}

func (aos *SubtractionOperationStrategy) Apply() error {
	aos.result = 0
	for i, arg := range aos.args {
		if i == 0 {
			aos.result = arg
		} else {
			aos.result -= arg
		}
	}
	aos.deductCostFromUserBalance()
	return nil
}

func NewSubtractionOperationStrategy() *SubtractionOperationStrategy {
	return &SubtractionOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      100,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   []float64{},
	}
}
