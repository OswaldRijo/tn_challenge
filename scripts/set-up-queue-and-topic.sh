#!/bin/sh
AWS_REGION=us-east-2
AWS_ACCOUNT=$(aws sts get-caller-identity --query "Account" --output text)
TOPIC_NAME=tn.user.created.topic
QUEUE_NAME=tn.user.created.queue

TOPIC_ARN=$(aws sns list-topics --query "Topics[?ends_with(TopicArn, ':$TOPIC_NAME')].TopicArn" --output text)
if [ -z "$TOPIC_ARN" ]; then
    echo "Topic $TOPIC_NAME does not exist. Creating..."
    TOPIC_ARN=$(aws sns create-topic --name $TOPIC_NAME --region $AWS_REGION --query "TopicArn" --output text)
    echo "Topic created. ARN: $TOPIC_ARN"
else
    echo "Topic $TOPIC_NAME already exists. ARN: $TOPIC_ARN"
fi

QUEUE_URL=$(aws sqs get-queue-url --queue-name "$QUEUE_NAME" --query "QueueUrl" --output text 2>/dev/null)

if [ -z "$QUEUE_URL" ]; then
    echo "Queue $QUEUE_NAME does not exists. Creating..."
    QUEUE_URL=$(aws sqs create-queue --queue-name "$QUEUE_NAME" --query "QueueUrl" --output text)
    echo "Queue created. URL: $QUEUE_URL"
    aws sns subscribe --topic-arn arn:aws:sns:$AWS_REGION:$AWS_ACCOUNT:$TOPIC_NAME --protocol sqs --notification-endpoint arn:aws:sqs:$AWS_REGION:$AWS_ACCOUNT:$QUEUE_NAME
else
    echo "Queue $QUEUE_NAME already exists. URL: $QUEUE_URL"
fi
