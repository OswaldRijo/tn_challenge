package operations_strategies

type RandStringOperationStrategy struct {
	*OperationStrategyImpl
	result string
	cost   float64
}

func (rsos *RandStringOperationStrategy) Apply() error {
	// TODO: "implement me"
	rsos.deductCostFromUserBalance()
	return nil
}

func (rsos *RandStringOperationStrategy) GetResultAsJson() ([]byte, error) {
	return serializeResultAsJson(rsos.result)
}

func (rsos *RandStringOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeResultAsJson([]byte{})
}

func NewRandStringOperationStrategy() *RandStringOperationStrategy {
	return &RandStringOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      100,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: "",
	}
}
