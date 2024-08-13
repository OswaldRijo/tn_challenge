package operations_strategies

type DivisionOperationStrategy struct {
	*OperationStrategyImpl
	result float64
	args   []float64
}

func (dos *DivisionOperationStrategy) setArgs(args ...float64) *DivisionOperationStrategy {
	dos.args = args
	return dos
}

func (dos *DivisionOperationStrategy) GetResultAsJson() ([]byte, error) {
	return serializeResultAsJson(dos.result)
}

func (dos *DivisionOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeResultAsJson(dos.args)
}

func (dos *DivisionOperationStrategy) Apply() error {
	dos.result = 0
	for i, arg := range dos.args {
		if i == 0 {
			dos.result = arg
		} else {
			dos.result /= arg
		}
	}
	dos.deductCostFromUserBalance()

	return nil
}

func NewDivisionOperationStrategy() *DivisionOperationStrategy {
	return &DivisionOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      100,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: float64(0),
		args:   []float64{},
	}
}
