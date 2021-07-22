// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package transformer

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
)

// Centralize anything related to ECS into a common file.
// TODO: Look at the fields to align them with ECS.
// TODO: how to keep the fields in sync with AWS?
// TODO: api gateway proxy a lot more information is available.

// CloudwatchLogs takes an CloudwatchLogsData and transform it into a beat event.
func CloudwatchLogs(request events.CloudwatchLogsData) []beat.Event {
	events := make([]beat.Event, len(request.LogEvents))

	for idx, logEvent := range request.LogEvents {
		events[idx] = beat.Event{
			Timestamp: time.Unix(0, logEvent.Timestamp*1000000),
			Fields: common.MapStr{
				"event": common.MapStr{
					"kind": "event",
				},
				"cloud": common.MapStr{
					"provider": "aws",
				},
				"message":              logEvent.Message,
				"id":                   logEvent.ID,
				"owner":                request.Owner,
				"log_stream":           request.LogStream,
				"log_group":            request.LogGroup,
				"message_type":         request.MessageType,
				"subscription_filters": request.SubscriptionFilters,
			},
		}
	}

	return events
}

// APIGatewayProxyRequest takes a web request on the api gateway proxy and transform it into a beat event.
func APIGatewayProxyRequest(request events.APIGatewayProxyRequest) beat.Event {
	return beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"event": common.MapStr{
				"kind":     "event",
				"category": []string{"network"},
				"type":     []string{"connection", "protocol"},
			},
			"cloud": common.MapStr{
				"provider":   "aws",
				"account.id": request.RequestContext.AccountID,
			},
			"network": common.MapStr{
				"transport": "tcp",
				"protocol":  "http",
			},
			"resource":          request.Resource,
			"path":              request.Path,
			"method":            request.HTTPMethod,
			"headers":           request.Headers,               // TODO: ECS map[string]
			"query_string":      request.QueryStringParameters, // TODO: map[string], might conflict with ECS
			"path_parameters":   request.PathParameters,
			"body":              request.Body, // TODO: could be JSON, json processor? could be used by other functions.
			"is_base64_encoded": request.IsBase64Encoded,
		},
	}
}

// KinesisEvent takes a kinesis event and create multiples beat events.
// DOCS: https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html
func KinesisEvent(request events.KinesisEvent) []beat.Event {
	events := make([]beat.Event, len(request.Records))
	for idx, record := range request.Records {
		events[idx] = beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"event": common.MapStr{
					"kind": "event",
				},
				"cloud": common.MapStr{
					"provider": "aws",
					"region":   record.AwsRegion,
				},
				"event_id":                record.EventID,
				"event_name":              record.EventName,
				"event_source":            record.EventSource,
				"event_source_arn":        record.EventSourceArn,
				"event_version":           record.EventVersion,
				"aws_region":              record.AwsRegion,
				"message":                 string(record.Kinesis.Data),
				"kinesis_partition_key":   record.Kinesis.PartitionKey,
				"kinesis_schema_version":  record.Kinesis.KinesisSchemaVersion,
				"kinesis_sequence_number": record.Kinesis.SequenceNumber,
				"kinesis_encryption_type": record.Kinesis.EncryptionType,
			},
		}
	}
	return events
}

// CloudwatchKinesisEvent takes a Kinesis event containing Cloudwatch logs and creates events for all
// Cloudwatch logs.
func CloudwatchKinesisEvent(request events.KinesisEvent, base64Encoded, compressed bool) ([]beat.Event, error) {
	var evts []beat.Event
	for _, record := range request.Records {
		envelopeFields := common.MapStr{
			"event": common.MapStr{
				"kind": "event",
			},
			"cloud": common.MapStr{
				"provider": "aws",
				"region":   record.AwsRegion,
			},
			"event_id":                record.EventID,
			"event_name":              record.EventName,
			"event_source":            record.EventSource,
			"event_source_arn":        record.EventSourceArn,
			"event_version":           record.EventVersion,
			"aws_region":              record.AwsRegion,
			"kinesis_partition_key":   record.Kinesis.PartitionKey,
			"kinesis_schema_version":  record.Kinesis.KinesisSchemaVersion,
			"kinesis_sequence_number": record.Kinesis.SequenceNumber,
			"kinesis_encryption_type": record.Kinesis.EncryptionType,
		}

		kinesisData := record.Kinesis.Data
		if base64Encoded {
			var err error
			kinesisData, err = base64.StdEncoding.DecodeString(string(kinesisData))
			if err != nil {
				return nil, err
			}
		}

		if compressed {
			inBuf := bytes.NewBuffer(record.Kinesis.Data)
			r, err := gzip.NewReader(inBuf)
			if err != nil {
				return nil, err
			}

			var outBuf bytes.Buffer
			_, err = io.Copy(&outBuf, r)
			if err != nil {
				r.Close()
				return nil, err
			}

			err = r.Close()
			if err != nil {
				return nil, err
			}
			kinesisData = outBuf.Bytes()
		}

		var cloudwatchEvents events.CloudwatchLogsData
		err := json.Unmarshal(kinesisData, &cloudwatchEvents)
		if err != nil {
			return nil, err
		}

		cwEvts := CloudwatchLogs(cloudwatchEvents)
		for _, cwe := range cwEvts {
			cwe.Fields.DeepUpdate(envelopeFields)
			evts = append(evts, cwe)
		}
	}
	return evts, nil
}

// SQS takes a SQS event and create multiples beat events.
func SQS(request events.SQSEvent) []beat.Event {
	events := make([]beat.Event, len(request.Records))
	for idx, record := range request.Records {
		events[idx] = beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"event": common.MapStr{
					"kind": "event",
				},
				"cloud": common.MapStr{
					"provider": "aws",
					"region":   record.AWSRegion,
				},
				"message_id":       record.MessageId,
				"receipt_handle":   record.ReceiptHandle,
				"message":          record.Body,
				"attributes":       record.Attributes,
				"event_source":     record.EventSource,
				"event_source_arn": record.EventSourceARN,
				"aws_region":       record.AWSRegion,
			},
			// TODO: SQS message attributes missing, need to check doc
		}
	}
	return events
}

func S3GetEvents(request events.S3Event) ([]beat.Event, error) {
	var evts []beat.Event
	svc := s3.New(session.New())
	for _, record := range request.Records {
		unescaped_key, err := url.QueryUnescape(record.S3.Object.Key)

		if err != nil {
			fmt.Println("Got error unescaping key: %s", record.S3.Object.Key)
			fmt.Println(err.Error())
			return nil, err
		}

		result, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(record.S3.Bucket.Name),
			Key:    aws.String(unescaped_key),
		})

		if err != nil {
			fmt.Println("Got error calling GetObject:")
			fmt.Println(err.Error())
			return nil, err
		}

		defer result.Body.Close()
		obj, err := ioutil.ReadAll(result.Body)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if obj[0] == 31 && obj[1] == 139 {
			inBuf := bytes.NewBuffer(obj)
			r, err := gzip.NewReader(inBuf)
			if err != nil {
				return nil, err
			}

			var outBuf bytes.Buffer
			_, err = io.Copy(&outBuf, r)
			if err != nil {
				r.Close()
				return nil, err
			}

			err = r.Close()
			if err != nil {
				return nil, err
			}
			obj = outBuf.Bytes()
		}

		obj_r := bytes.NewReader(obj)
		obj_line := bufio.NewScanner(obj_r)

		for obj_line.Scan() {
			s3evt := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"event": common.MapStr{
						"kind": "event",
					},
					"cloud": common.MapStr{
						"provider": "aws",
						"region":   record.AWSRegion,
					},
					"message":      obj_line.Text(),
					"event_source": record.EventSource,
					"bucket_name":  record.S3.Bucket.Name,
					"bucket_key":   unescaped_key,
					"aws_region":   record.AWSRegion,
				},
			}
			evts = append(evts, s3evt)
		}
	}
	return evts, nil
}
