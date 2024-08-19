# Project Challenge

This project is part of a technical challenge aimed at assessing programming skills. The project involves setting up a local environment, running scripts, and interacting with AWS services.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Init Setup](#init-setup)
- [Run App](#run-app)
- [Kill App](#kill-app)
- [Troubleshooting](#troubleshooting)

## Prerequisites

Before setting up the environment, certain dependencies need to be installed based on your operating system:

- **For MacOS users**:
    - Ensure **Go** and **Node** are both installed   
    - Ensure that you have **Homebrew** and **AWS CLI** installed.
    - Ensure **make** command is available
    - Ensure **docker-compose** command is available

- **For Linux (Ubuntu) users**:
    - Ensure **Go** and **Node** are both installed
    - Ensure that you have **AWS CLI** installed.
    - Ensure **make** command is available
    - Ensure **docker compose** command is available

**Note:**
- The installation script may require root permissions.
- You must be logged into AWS in the terminal where you will run the setup, as the installation script creates resources in AWS.

## Init Setup

To install the required dependencies, run the following command:

```bash
./scripts/setup-local-env.sh
```

After the script has successfully run, ensure that the QUEUE is able to consume messages from the SNS topic.
To do so:
1. Go to AWS Console.
2. Select us-region-2, 
3. Go to [Simple Notification Service](https://us-east-2.console.aws.amazon.com/sns/v3/home?region=us-east-2#/topics)
4. Select tn_user_created_topic
5. Click on `Publish Messaje` button
6. Type message of your preference, click on `Publish Message`.

After that go to [Simple Queue Service](https://us-east-2.console.aws.amazon.com/sqs/v3/home?region=us-east-2#/queues) and check if the message was received by the queue.
To do so:

1. Click on `tn_user_created_queue`
2. Click on `Send and receive messages`
3. Click on `Poll for messages`
4. Et voila. You should be able to see the message you just sent. 

Make sure you remove the message to prevent a failure within the app.
If you're not seeing the message please go to [Troubleshooting section](#troubleshooting)


## Run App

To start the application, execute the following command:
```bash
./scripts/run-app.sh
```

## Kill App

To stop the application, run the following command:
```bash
./scripts/kill-app.sh
```
The command will remove all docker images previously created but the database image, reason why i encourage you to check it and remove it by yourself

## Troubleshooting
### AWS CLI: "Unable to locate credentials"

This error is related to AWS authentication. To resolve this issue:

1. Refresh your AWS credentials.
2. Re-run the script that failed.

### SNS/SQS: Issues with message consumption in AWS region `us-east-2`
If the **QUEUE** is not consuming messages from the **SNS** topic, follow these steps:

1. Open the **AWS SNS Console** and locate the topic named `tn_user_created_topic`.
2. Delete the existing subscription to the queue `tn_user_created_queue`.
3. Recreate the subscription via the **AWS UI**.
4. Once the subscription is re-established, publish a test message from the **AWS SNS UI** to verify that the queue is correctly consuming messages.

If the **QUEUE** is still not consuming messages from the **SNS** topic, follow these steps:

1. Open the **AWS SQS Console** and locate the queue `tn_user_created_queue`
2. Delete the existing subscription topic named `tn_user_created_topic`.
3. Recreate the subscription via the **AWS UI** this time from SQS service.
4. Once the subscription is re-established, publish a test message from the **AWS SNS UI** to verify that the queue is correctly consuming messages.