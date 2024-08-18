package publisher

import (
	"context"
)

type ProducerImpl struct {
	topicArn string
}

//go:generate mockery --name=Producer --output=../../../mocks/packages/pubsub
type Producer interface {
	SendMessage(ctx context.Context, message map[string]any) error
}

func NewProducer(topicArn string) Producer {
	return &ProducerImpl{
		topicArn: topicArn,
	}
}
