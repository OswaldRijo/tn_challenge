package consumer

import (
	"context"
	"encoding/json"

	awsSqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	"truenorth/packages/logger"
)

func (c *ConsumerImpl) getQueueURL(queue *string) (*awsSqs.GetQueueUrlOutput, error) {
	urlResult, err := c.awsClient.GetQueueUrl(context.TODO(), &awsSqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}

	return urlResult, nil
}

func (c *ConsumerImpl) getMessages(ctx context.Context, queueURL *string) (*awsSqs.ReceiveMessageOutput, error) {
	msgResult, err := c.awsClient.ReceiveMessage(ctx, &awsSqs.ReceiveMessageInput{
		AttributeNames: []types.QueueAttributeName{
			types.QueueAttributeName(types.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   10,
	})
	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func (c *ConsumerImpl) AddQueue(queueName string, controller func(ctx context.Context, message map[string]any) error) {
	c.consumers[queueName] = controller
}

func (c *ConsumerImpl) Start() {

	for k, controller := range c.consumers {
		queue, err := c.getQueueURL(&k)

		if err != nil {
			logger.GetLogger().Panic(context.TODO(), "err", err)
		}

		controller1 := controller
		go func() {
			for {
				ctx := context.TODO()
				output, err := c.getMessages(ctx, queue.QueueUrl)

				if err != nil {
					logger.GetLogger().Panic(context.TODO(), "err", err)
				}

				if len(output.Messages) > 0 {
					var message map[string]any
					var body map[string]any
					_ = json.Unmarshal([]byte(*output.Messages[0].Body), &message)
					_ = json.Unmarshal([]byte(message["Message"].(string)), &body)
					err = controller1(ctx, body)
					if err != nil {
						logger.GetLogger().Error(context.TODO(), "err", err)
						continue
					}
					err = c.deleteMessages(ctx, queue.QueueUrl, output.Messages[0].ReceiptHandle)
					if err != nil {
						logger.GetLogger().Error(context.TODO(), "err", err)
					}
				}
			}
		}()
	}

}

func (c *ConsumerImpl) deleteMessages(ctx context.Context, queueURL *string, receiptHandle *string) error {

	_, err := c.awsClient.DeleteMessage(ctx, &awsSqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: receiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
