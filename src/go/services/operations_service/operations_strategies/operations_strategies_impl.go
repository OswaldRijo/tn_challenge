package operations_strategies

import (
	"encoding/json"
	"log"

	operationspb "truenorth/pb/operations"
)

type OperationStrategy interface {
	Apply() error
	GetArgsAsJson() ([]byte, error)
	GetResultAsJson() ([]byte, error)
	GetResultantUserBalance() float64
	GetCost() float64
}

type OperationStrategyImpl struct {
	cost                      float64
	userBalance               float64
	userBalanceAfterOperation float64
}

func (osi *OperationStrategyImpl) setUserBalance(userBalance float64) {
	osi.userBalance = userBalance
}

func (osi *OperationStrategyImpl) deductCostFromUserBalance() {
	osi.userBalanceAfterOperation = osi.userBalance - osi.cost
}

func (osi *OperationStrategyImpl) GetResultantUserBalance() float64 {
	return osi.userBalanceAfterOperation
}

func (osi *OperationStrategyImpl) GetCost() float64 {
	return osi.cost
}

func serializeResultAsJson[T any](result T) ([]byte, error) {
	data := map[string]T{"result": result}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error serializing result as JSON: %v", err)
		return nil, err
	}
	return jsonData, nil
}

func serializeArgsAsJson[T any](args ...T) ([]byte, error) {
	data := map[string][]T{"args": args}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error serializing result as JSON: %v", err)
		return nil, err
	}
	return jsonData, nil
}

func NewOperationStrategy(opType operationspb.OperationType, userBalance float64, args ...float64) OperationStrategy {
	switch opType {
	case operationspb.OperationType_SUBTRACTION:
		strategy := NewSubtractionOperationStrategy().setArgs(args...)
		strategy.setUserBalance(userBalance)
		return strategy
	case operationspb.OperationType_MULTIPLICATION:
		strategy := NewMultiplicationOperationStrategy().setArgs(args...)
		strategy.setUserBalance(userBalance)
		return strategy
	case operationspb.OperationType_DIVISION:
		strategy := NewDivisionOperationStrategy().setArgs(args...)
		strategy.setUserBalance(userBalance)
		return strategy
	case operationspb.OperationType_RANDOM_STRING:
		strategy := NewRandStringOperationStrategy()
		strategy.setUserBalance(userBalance)
		return strategy
	default:
		strategy := NewMultiplicationOperationStrategy().setArgs(args...)
		strategy.setUserBalance(userBalance)
		return strategy
	}
}
