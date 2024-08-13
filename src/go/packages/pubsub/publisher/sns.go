package publisher

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	awsSns "github.com/aws/aws-sdk-go-v2/service/sns"

	"truenorth/packages/utils"
)

func (p *ProducerImpl) SendMessage(ctx context.Context, message map[string]any) error {
	cfg, err := config.LoadDefaultConfig(ctx)

	client := awsSns.NewFromConfig(cfg)
	messageStr, _ := json.Marshal(message)

	input := &awsSns.PublishInput{
		Message:  utils.Pointer(string(messageStr)),
		TopicArn: &p.topicArn,
	}

	_, err = client.Publish(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
