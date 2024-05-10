// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !aix

package azureeventhub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/checkpoints"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	v2 "github.com/elastic/beats/v7/filebeat/input/v2"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"time"
)

type eventHubInputV2 struct {
	config          azureInputConfig
	log             *logp.Logger
	metrics         *inputMetrics
	checkpointStore *checkpoints.BlobStore
	consumerClient  *azeventhubs.ConsumerClient
	client          beat.Client
}

func newEventHubInputV2(config azureInputConfig, log *logp.Logger) (v2.Input, error) {
	return &eventHubInputV2{
		config: config,
		log:    log.Named(inputName),
	}, nil
}

func (in *eventHubInputV2) Name() string {
	return inputName
}

func (in *eventHubInputV2) Test(v2.TestContext) error {
	return nil
}

func (in *eventHubInputV2) Run(
	inputContext v2.Context,
	pipeline beat.Pipeline,
) error {
	var err error

	// Create client for publishing events and receive notification of their ACKs.
	in.client, err = createPipelineClient(pipeline)
	if err != nil {
		return fmt.Errorf("failed to create pipeline client: %w", err)
	}
	defer in.client.Close()

	// Setup input metrics
	inputMetrics := newInputMetrics(inputContext.ID, nil)
	if err != nil {
		return fmt.Errorf("failed to create input metrics: %w", err)
	}
	defer inputMetrics.Close()
	in.metrics = inputMetrics

	ctx := v2.GoContextFromCanceler(inputContext.Cancelation)

	// Initialize everything for this run
	err = in.setup(ctx)
	if err != nil {
		return err
	}
	defer in.consumerClient.Close(context.Background())

	// Start the main run loop
	in.run(ctx)

	return nil
}

func (in *eventHubInputV2) setup(ctx context.Context) error {
	// FIXME: check more client creation options.
	blobContainerClient, err := container.NewClientFromConnectionString(
		in.config.SAConnectionString,
		in.config.SAContainer,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create blob container client: %w", err)
	}

	checkpointStore, err := checkpoints.NewBlobStore(blobContainerClient, nil)
	if err != nil {
		return fmt.Errorf("failed to create checkpoint store: %w", err)
	}
	in.checkpointStore = checkpointStore

	consumerClient, err := azeventhubs.NewConsumerClientFromConnectionString(
		in.config.ConnectionString,
		in.config.EventHubName,
		in.config.ConsumerGroup,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create consumer client: %w", err)
	}
	in.consumerClient = consumerClient

	return nil
}

func (in *eventHubInputV2) run(ctx context.Context) {
	processor, err := azeventhubs.NewProcessor(
		in.consumerClient,
		in.checkpointStore,
		nil,
	)
	if err != nil {
		in.log.Errorw("error creating event processor", "error", err)
		return
	}

	// Run in the background, launching goroutines to process each partition
	go in.workersLoop(processor)

	if err := processor.Run(ctx); err != nil {
		in.log.Errorw("error running processor", "error", err)
	}
}

func (in *eventHubInputV2) workersLoop(processor *azeventhubs.Processor) {
	for {
		processorPartitionClient := processor.NextPartitionClient(context.TODO())

		if processorPartitionClient == nil {
			// Processor has stopped
			break
		}

		go func() {
			if err := in.processEventsForPartition(processorPartitionClient); err != nil {
				//panic(err)
				logp.Info("error processing events for partition: %v", err)
			}
		}()
	}
}

// processEventsForPartition shows the typical pattern for processing a partition.
func (in *eventHubInputV2) processEventsForPartition(partitionClient *azeventhubs.ProcessorPartitionClient) error {
	// 1. [BEGIN] Initialize any partition specific resources for your application.
	// 2. [CONTINUOUS] Loop, calling ReceiveEvents() and UpdateCheckpoint().
	// 3. [END] Cleanup any resources.

	defer func() {
		// 3/3 [END] Do cleanup here, like shutting down database clients
		// or other resources used for processing this partition.
		shutdownPartitionResources(partitionClient)
	}()

	// 1/3 [BEGIN] Initialize any partition specific resources for your application.
	if err := initializePartitionResources(partitionClient.PartitionID()); err != nil {
		return err
	}

	// 2/3 [CONTINUOUS] Receive events, checkpointing as needed using UpdateCheckpoint.
	for {
		// Wait up to a minute for 100 events, otherwise returns whatever we collected during that time.
		receiveCtx, cancelReceive := context.WithTimeout(context.TODO(), 10*time.Second)
		events, err := partitionClient.ReceiveEvents(receiveCtx, 100, nil)
		cancelReceive()

		if err != nil && !errors.Is(err, context.DeadlineExceeded) {
			var eventHubError *azeventhubs.Error

			if errors.As(err, &eventHubError) && eventHubError.Code == azeventhubs.ErrorCodeOwnershipLost {
				return nil
			}

			return err
		}

		if len(events) == 0 {
			continue
		}

		err = in.processReceivedEvents(events)
		if err != nil {
			return fmt.Errorf("error processing received events: %w", err)
		}

		// Updates the checkpoint with the latest event received. If processing needs to restart
		// it will restart from this point, automatically.
		if err := partitionClient.UpdateCheckpoint(context.TODO(), events[len(events)-1], nil); err != nil {
			return err
		}
	}
}

func (in *eventHubInputV2) processReceivedEvents(receivedEvents []*azeventhubs.ReceivedEventData) error {
	processingStartTime := time.Now()
	azure := mapstr.M{
		// The partition ID is not available.
		// "partition_id":   partitionID,
		"eventhub":       in.config.EventHubName,
		"consumer_group": in.config.ConsumerGroup,
	}

	for _, receivedEventData := range receivedEvents {
		// A single event can contain multiple records. We create a new event for each record.
		records := in.parseEvent(receivedEventData.Body)

		for record := range records {
			_, _ = azure.Put("offset", receivedEventData.Offset)
			_, _ = azure.Put("sequence_number", receivedEventData.SequenceNumber)
			_, _ = azure.Put("enqueued_time", receivedEventData.EnqueuedTime)

			event := beat.Event{
				// this is the default value for the @timestamp field; usually the ingest
				// pipeline replaces it with a value in the payload.
				Timestamp: processingStartTime,
				Fields: mapstr.M{
					"message": record,
					"azure":   azure,
				},
				Private: receivedEventData.Body,
			}

			in.client.Publish(event)
		}
	}

	return nil
}

func (in *eventHubInputV2) parseEvent(bMessage []byte) []string {
	var mapObject map[string][]interface{}
	var records []string

	// Clean up the message for known issues [1] where Azure services produce malformed JSON documents.
	// Sanitization occurs if options are available and the message contains an invalid JSON.
	//
	// [1]: https://learn.microsoft.com/en-us/answers/questions/1001797/invalid-json-logs-produced-for-function-apps
	if len(in.config.SanitizeOptions) != 0 && !json.Valid(bMessage) {
		bMessage = sanitize(bMessage, in.config.SanitizeOptions...)
		in.metrics.sanitizedMessages.Inc()
	}

	// check if the message is a "records" object containing a list of events
	err := json.Unmarshal(bMessage, &mapObject)
	if err == nil {
		if len(mapObject[expandEventListFromField]) > 0 {
			for _, ms := range mapObject[expandEventListFromField] {
				js, err := json.Marshal(ms)
				if err == nil {
					records = append(records, string(js))
					in.metrics.receivedEvents.Inc()
				} else {
					in.log.Errorw(fmt.Sprintf("serializing message %s", ms), "error", err)
				}
			}
		}
	} else {
		in.log.Debugf("deserializing multiple messages to a `records` object returning error: %s", err)
		// in some cases the message is an array
		var arrayObject []interface{}
		err = json.Unmarshal(bMessage, &arrayObject)
		if err != nil {
			// return entire message
			in.log.Debugf("deserializing multiple messages to an array returning error: %s", err)
			in.metrics.decodeErrors.Inc()
			return []string{string(bMessage)}
		}

		for _, ms := range arrayObject {
			js, err := json.Marshal(ms)
			if err == nil {
				records = append(records, string(js))
				in.metrics.receivedEvents.Inc()
			} else {
				in.log.Errorw(fmt.Sprintf("serializing message %s", ms), "error", err)
			}
		}
	}

	return records
}

func initializePartitionResources(partitionID string) error {
	// initialize things that might be partition specific, like a
	// database connection.
	return nil
}

func shutdownPartitionResources(partitionClient *azeventhubs.ProcessorPartitionClient) {
	// Each PartitionClient holds onto an external resource and should be closed if you're
	// not processing them anymore.
	defer partitionClient.Close(context.TODO())
}
