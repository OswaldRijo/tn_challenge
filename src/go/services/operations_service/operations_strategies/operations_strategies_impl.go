package operations_strategies

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/context"

	operationspb "truenorth/pb/operations"
)

type OperationStrategy interface {
	Apply(ctx context.Context) error
	GetArgsAsJson() ([]byte, error)
	GetResult() string
	GetResultantUserBalance() float64
	GetCost() float64
	setUserBalance(balance float64)
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

func parseResultToString[T any](result T) string {
	return fmt.Sprintf("%v", result)
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
	var strategy OperationStrategy
	switch opType {
	case operationspb.OperationType_SUBTRACTION:
		strategy = NewSubtractionOperationStrategy(args...)
		break
	case operationspb.OperationType_MULTIPLICATION:
		strategy = NewMultiplicationOperationStrategy(args...)
		break
	case operationspb.OperationType_DIVISION:
		strategy = NewDivisionOperationStrategy(args...)
		break
	case operationspb.OperationType_RANDOM_STRING:
		strategy = NewRandStringOperationStrategy()
		break
	case operationspb.OperationType_SQUARE_ROOT:
		strategy = NewSquareRootOperationStrategy(args...)
		break
	default:
		strategy = NewAdditionOperationStrategy(args...)
		break
	}
	strategy.setUserBalance(userBalance)
	return strategy
}
