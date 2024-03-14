// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package awss3

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/elastic/elastic-agent-libs/logp"
)

const testTimeout = 10 * time.Second

var (
	errFakeConnectivityFailure = errors.New("fake connectivity failure")
	errFakeGetAttributeFailute = errors.New("something went wrong")
)

func TestSQSReceiver(t *testing.T) {
	logp.TestingSetup()

	const maxMessages = 5

	t.Run("ReceiveMessage success", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
		defer cancel()

		ctrl, ctx := gomock.WithContext(ctx, t)
		defer ctrl.Finish()
		mockAPI := NewMockSQSAPI(ctrl)
		mockMsgHandler := NewMockSQSProcessor(ctrl)
		msg := newSQSMessage(newS3Event("log.json"))

		gomock.InOrder(
			// Initial ReceiveMessage for maxMessages.
			mockAPI.EXPECT().
				ReceiveMessage(gomock.Any(), gomock.Eq(maxMessages)).
				Times(1).
				DoAndReturn(func(_ context.Context, _ int) ([]types.Message, error) {
					// Return single message.
					return []types.Message{msg}, nil
				}),

			// Follow up ReceiveMessages for either maxMessages-1 or maxMessages
			// depending on how long processing of previous message takes.
			mockAPI.EXPECT().
				ReceiveMessage(gomock.Any(), gomock.Any()).
				Times(1).
				DoAndReturn(func(_ context.Context, _ int) ([]types.Message, error) {
					// Stop the test.
					cancel()
					return nil, nil
				}),
		)

		mockClient := NewMockBeatClient(ctrl)
		mockBeatPipeline := NewMockBeatPipeline(ctrl)

		mockBeatPipeline.EXPECT().ConnectWith(gomock.Any()).Return(mockClient, nil).Times(maxMessages)

		// Expect the one message returned to have been processed.
		mockMsgHandler.EXPECT().
			ProcessSQS(gomock.Any(), gomock.Eq(&msg), gomock.Any(), gomock.Any(), gomock.Any()).
			Times(1).
			DoAndReturn(
				func(ctx context.Context, msg *types.Message, _ beat.Client, acker *EventACKTracker, _ time.Time) (uint64, error) {
					_, keepaliveCancel := context.WithCancel(ctx)
					log := log.Named("sqs_s3_event")
					acker.MarkSQSProcessedWithData(msg, 1, -1, time.Now(), nil, nil, keepaliveCancel, new(sync.WaitGroup), mockMsgHandler, log)
					acker.ACK()
					<-acker.ctx.Done()
					return 1, nil
				})

		// Expect the client to be closed
		mockClient.EXPECT().Close().Times(maxMessages)

		// Expect the one message returned to have been deleted.
		mockMsgHandler.EXPECT().
			DeleteSQS(gomock.Eq(&msg), gomock.Any(), gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil)

		// Execute sqsReader and verify calls/state.
		receiver := newSQSReader(logp.NewLogger(inputName), nil, mockAPI, maxMessages, mockMsgHandler, mockBeatPipeline)
		require.NoError(t, receiver.Receive(ctx))
		assert.Equal(t, maxMessages, receiver.workerSem.Available())
	})

	t.Run("retry after ReceiveMessage error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), sqsRetryDelay+testTimeout)
		defer cancel()

		ctrl, ctx := gomock.WithContext(ctx, t)
		defer ctrl.Finish()
		mockAPI := NewMockSQSAPI(ctrl)
		mockMsgHandler := NewMockSQSProcessor(ctrl)

		mockClient := NewMockBeatClient(ctrl)
		mockBeatPipeline := NewMockBeatPipeline(ctrl)

		mockBeatPipeline.EXPECT().ConnectWith(gomock.Any()).Return(mockClient, nil).Times(maxMessages)

		gomock.InOrder(
			// Initial ReceiveMessage gets an error.
			mockAPI.EXPECT().
				ReceiveMessage(gomock.Any(), gomock.Eq(maxMessages)).
				Times(1).
				DoAndReturn(func(_ context.Context, _ int) ([]types.Message, error) {
					return nil, errFakeConnectivityFailure
				}),
			// After waiting for sqsRetryDelay, it retries.
			mockAPI.EXPECT().
				ReceiveMessage(gomock.Any(), gomock.Eq(maxMessages)).
				Times(1).
				DoAndReturn(func(_ context.Context, _ int) ([]types.Message, error) {
					cancel()
					return nil, nil
				}),
		)

		// Expect the client to be closed
		mockClient.EXPECT().Close().Times(maxMessages)

		// Execute SQSReceiver and verify calls/state.
		receiver := newSQSReader(logp.NewLogger(inputName), nil, mockAPI, maxMessages, mockMsgHandler, mockBeatPipeline)
		require.NoError(t, receiver.Receive(ctx))
		assert.Equal(t, maxMessages, receiver.workerSem.Available())
	})
}

func TestGetApproximateMessageCount(t *testing.T) {
	logp.TestingSetup()

	const maxMessages = 5
	const count = 500
	attrName := []types.QueueAttributeName{sqsApproximateNumberOfMessages}
	attr := map[string]string{"ApproximateNumberOfMessages": "500"}

	t.Run("GetApproximateMessageCount success", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
		defer cancel()

		ctrl, ctx := gomock.WithContext(ctx, t)
		defer ctrl.Finish()
		mockAPI := NewMockSQSAPI(ctrl)
		mockMsgHandler := NewMockSQSProcessor(ctrl)
		mockBeatPipeline := NewMockBeatPipeline(ctrl)

		gomock.InOrder(
			mockAPI.EXPECT().
				GetQueueAttributes(gomock.Any(), gomock.Eq(attrName)).
				Times(1).
				DoAndReturn(func(_ context.Context, _ []types.QueueAttributeName) (map[string]string, error) {
					return attr, nil
				}),
		)

		receiver := newSQSReader(logp.NewLogger(inputName), nil, mockAPI, maxMessages, mockMsgHandler, mockBeatPipeline)
		receivedCount, err := receiver.GetApproximateMessageCount(ctx)
		assert.Equal(t, count, receivedCount)
		assert.Nil(t, err)
	})

	t.Run("GetApproximateMessageCount error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
		defer cancel()

		ctrl, ctx := gomock.WithContext(ctx, t)
		defer ctrl.Finish()

		mockAPI := NewMockSQSAPI(ctrl)
		mockMsgHandler := NewMockSQSProcessor(ctrl)
		mockBeatPipeline := NewMockBeatPipeline(ctrl)

		gomock.InOrder(
			mockAPI.EXPECT().
				GetQueueAttributes(gomock.Any(), gomock.Eq(attrName)).
				Times(1).
				DoAndReturn(func(_ context.Context, _ []types.QueueAttributeName) (map[string]string, error) {
					return nil, errFakeGetAttributeFailute
				}),
		)

		receiver := newSQSReader(logp.NewLogger(inputName), nil, mockAPI, maxMessages, mockMsgHandler, mockBeatPipeline)
		receivedCount, err := receiver.GetApproximateMessageCount(ctx)
		assert.Equal(t, -1, receivedCount)
		assert.NotNil(t, err)
	})
}

func newSQSMessage(events ...s3EventV2) types.Message {
	body, err := json.Marshal(s3EventsV2{Records: events})
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(body)
	id, _ := uuid.FromBytes(hash[:16])
	messageID := id.String()
	receipt := "receipt-" + messageID
	bodyStr := string(body)

	return types.Message{
		Body:          &bodyStr,
		MessageId:     &messageID,
		ReceiptHandle: &receipt,
	}
}

func newSNSSQSMessage() types.Message {
	body, err := json.Marshal(s3EventsV2{
		TopicArn: "arn:aws:sns:us-east-1:1234:sns-topic",
		Message:  "{\"Records\":[{\"eventSource\":\"aws:s3\",\"awsRegion\":\"us-east-1\",\"eventName\":\"ObjectCreated:Put\",\"s3\":{\"configurationId\":\"sns-notification-vpc-flow-logs\",\"bucket\":{\"name\":\"vpc-flow-logs-ks\",\"arn\":\"arn:aws:s3:::vpc-flow-logs-ks\"},\"object\":{\"key\":\"test-object-key\"}}}]}",
	})
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(body)
	id, _ := uuid.FromBytes(hash[:16])
	messageID := id.String()
	receipt := "receipt-" + messageID
	bodyStr := string(body)

	return types.Message{
		Body:          &bodyStr,
		MessageId:     &messageID,
		ReceiptHandle: &receipt,
	}
}

func newS3Event(key string) s3EventV2 {
	record := s3EventV2{
		AWSRegion:   "us-east-1",
		EventSource: "aws:s3",
		EventName:   "ObjectCreated:Put",
		Provider:    "aws",
	}
	record.S3.Bucket.Name = "foo"
	record.S3.Bucket.ARN = "arn:aws:s3:::foo"
	record.S3.Object.Key = key
	return record
}
