package consumer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	awsSqs "github.com/aws/aws-sdk-go-v2/service/sqs"
)

type ConsumerImpl struct {
	consumers map[string]func(ctx context.Context, message map[string]any) error
	awsClient *awsSqs.Client
}

type Consumer interface {
	getQueueURL(queue *string) (*awsSqs.GetQueueUrlOutput, error)
	getMessages(ctx context.Context, queueURL *string) (*awsSqs.ReceiveMessageOutput, error)
	AddQueue(queueName string, controller func(ctx context.Context, message map[string]any) error)
	Start()
}

func NewConsumer() Consumer {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := awsSqs.NewFromConfig(cfg)
	return &ConsumerImpl{
		consumers: make(map[string]func(ctx context.Context, message map[string]any) error),
		awsClient: client,
	}
}
