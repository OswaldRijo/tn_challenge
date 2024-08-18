#!/bin/sh

if command -v pnpm &> /dev/null; then
    echo "pnpm is already present. Version: $(pnpm --version)"
else
    echo "pnpm is not present. Downloading and instaling pnpm"

    if ! command -v node &> /dev/null; then
        echo "Node.js is not installed. Please, install Node.js first."
        exit 1
    fi

    npm install -g pnpm

    if command -v pnpm &> /dev/null; then
        echo "pnpm has been installed successfully. Version: $(pnpm --version)"
    else
        echo "pnpm not installed, please handle its installation manually."
    fi
fi

./scripts/install-mockery.sh
./scripts/install-buf.sh

sudo make protos_go
sudo make protos_npm

./scripts/install-docker-compose.sh
./scripts/install-psql.sh


AWS_REGION=us-east-2
AWS_ACCOUNT=$(aws sts get-caller-identity --query "Account" --output text)
TOPIC_NAME=tn_user_created_topic
QUEUE_NAME=tn_user_created_queue

TOPIC_ARN=$(aws sns --region $AWS_REGION list-topics --query "Topics[?ends_with(TopicArn, ':$TOPIC_NAME')].TopicArn" --output text)
if [ -z "$TOPIC_ARN" ]; then
    echo "Topic $TOPIC_NAME does not exist. Creating..."
    TOPIC_ARN=$(aws sns --region $AWS_REGION create-topic --name $TOPIC_NAME --region $AWS_REGION --query "TopicArn" --output text)
    echo "Topic created. ARN: $TOPIC_ARN"
else
    echo "Topic $TOPIC_NAME already exists. ARN: $TOPIC_ARN"
fi

QUEUE_URL=$(aws sqs --region $AWS_REGION get-queue-url --queue-name "$QUEUE_NAME" --query "QueueUrl" --output text 2>/dev/null)

if [ -z "$QUEUE_URL" ]; then
    echo "Queue $QUEUE_NAME does not exists. Creating..."
    QUEUE_URL=$(aws sqs --region $AWS_REGION create-queue --queue-name "$QUEUE_NAME" --query "QueueUrl" --output text)
    echo "Queue created. URL: $QUEUE_URL"
    aws sns subscribe --region $AWS_REGION --topic-arn arn:aws:sns:$AWS_REGION:$AWS_ACCOUNT:$TOPIC_NAME --protocol sqs --notification-endpoint arn:aws:sqs:$AWS_REGION:$AWS_ACCOUNT:$QUEUE_NAME --output text
else
    echo "Queue $QUEUE_NAME already exists. URL: $QUEUE_URL"
fi

touch .properties
echo USER_CREATED_QUEUE=$QUEUE_NAME > .properties
echo USER_CREATED_TOPIC_ARN=$TOPIC_ARN >> .properties
echo AWS_REGION=$AWS_REGION >> .properties
