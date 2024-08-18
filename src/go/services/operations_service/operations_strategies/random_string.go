package operations_strategies

import (
	"context"
	"io"
	"net/http"

	"truenorth/packages/logger"
	"truenorth/services/operations_service/config"
)

type RandStringOperationStrategy struct {
	*OperationStrategyImpl
	result string
	cost   float64
}

const RandStringsUrl = "https://www.random.org/strings/?num=1&len=32&digits=on&upperalpha=on&loweralpha=on&unique=on&format=plain&rnd=new"

func (rsos *RandStringOperationStrategy) Apply(ctx context.Context) error {
	res, err := http.Get(RandStringsUrl)
	if err != nil {
		logger.GetLog().Infof(ctx, "error making http request: %s\n", err)
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		logger.GetLog().Infof(ctx, "client: could not read response body: %s\n", err)
		return err
	}

	rsos.result = string(resBody)
	rsos.result = rsos.result[:len(rsos.result)-2]
	return rsos.deductCostFromUserBalance()
}

func (rsos *RandStringOperationStrategy) GetResult() string {
	return parseResultToString(rsos.result)
}

func (rsos *RandStringOperationStrategy) GetArgsAsJson() ([]byte, error) {
	return serializeArgsAsJson([]byte{})
}

func NewRandStringOperationStrategy() *RandStringOperationStrategy {
	return &RandStringOperationStrategy{
		OperationStrategyImpl: &OperationStrategyImpl{
			cost:                      config.Config.RandomStringOperationCost,
			userBalance:               0,
			userBalanceAfterOperation: 0,
		},
		result: "",
	}
}
