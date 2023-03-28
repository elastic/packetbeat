// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package awss3

import (
	"context"
	"io"
	"sync/atomic"
	"time"

	"github.com/rcrowley/go-metrics"

	"github.com/elastic/beats/v7/libbeat/monitoring/inputmon"
	"github.com/elastic/elastic-agent-libs/monitoring"
	"github.com/elastic/elastic-agent-libs/monitoring/adapter"
)

type inputMetrics struct {
	registry   *monitoring.Registry
	unregister func()
	ctx        context.Context
	cancel     context.CancelFunc

	utilizationNanos int64

	sqsMessagesReceivedTotal            *monitoring.Uint  // Number of SQS messages received (not necessarily processed fully).
	sqsVisibilityTimeoutExtensionsTotal *monitoring.Uint  // Number of SQS visibility timeout extensions.
	sqsMessagesInflight                 *monitoring.Uint  // Number of SQS messages inflight (gauge).
	sqsMessagesReturnedTotal            *monitoring.Uint  // Number of SQS message returned to queue (happens on errors implicitly after visibility timeout passes).
	sqsMessagesDeletedTotal             *monitoring.Uint  // Number of SQS messages deleted.
	sqsMessagesWaiting                  *monitoring.Int   // Number of SQS messages waiting in the SQS queue (gauge). The value is refreshed every minute via data from GetQueueAttributes.
	sqsWorkerUtilization                *monitoring.Float // Rate of SQS worker utilization over previous 5 seconds. 0 indicates idle, 1 indicates all workers utilized.
	sqsMessageProcessingTime            metrics.Sample    // Histogram of the elapsed SQS processing times in nanoseconds (time of receipt to time of delete/return).
	sqsLagTime                          metrics.Sample    // Histogram of the difference between the SQS SentTimestamp attribute and the time when the SQS message was received expressed in nanoseconds.

	s3ObjectsRequestedTotal *monitoring.Uint // Number of S3 objects downloaded.
	s3ObjectsAckedTotal     *monitoring.Uint // Number of S3 objects processed that were fully ACKed.
	s3ObjectsListedTotal    *monitoring.Uint // Number of S3 objects returned by list operations.
	s3ObjectsProcessedTotal *monitoring.Uint // Number of S3 objects that matched file_selectors rules.
	s3BytesProcessedTotal   *monitoring.Uint // Number of S3 bytes processed.
	s3EventsCreatedTotal    *monitoring.Uint // Number of events created from processing S3 data.
	s3ObjectsInflight       *monitoring.Uint // Number of S3 objects inflight (gauge).
	s3ObjectProcessingTime  metrics.Sample   // Histogram of the elapsed S3 object processing times in nanoseconds (start of download to completion of parsing).
}

// Close cancels the context and removes the metrics from the registry.
func (m *inputMetrics) Close() {
	m.cancel()
	m.unregister()
}

func (m *inputMetrics) updateSQSProcessingTime(d time.Duration) {
	m.sqsMessageProcessingTime.Update(d.Nanoseconds())
	atomic.AddInt64(&m.utilizationNanos, d.Nanoseconds())
}

func (m *inputMetrics) setSQSMessagesWaiting(count int64) {
	if m.sqsMessagesWaiting == nil {
		// if metric not initialized, and count is -1, do nothing
		if count == -1 {
			return
		}
		m.sqsMessagesWaiting = monitoring.NewInt(m.registry, "sqs_messages_waiting_gauge")
	}

	m.sqsMessagesWaiting.Set(count)
}

func calculateUtilizationAndReset(d time.Duration, maxMessagesInflight int, m *inputMetrics) float64 {
	maxUtilization := float64(d) * float64(maxMessagesInflight)
	utilizedRate := float64(atomic.SwapInt64(&m.utilizationNanos, 0)) / maxUtilization

	if utilizedRate == 0 {
		// Falling back to inflight stats when the workers are long running
		inflight := m.sqsMessagesInflight.Get()
		if inflight > 0 {
			return float64(inflight) / float64(maxMessagesInflight)
		}
	} else if utilizedRate > 1 {
		// Normalizing the utilization after long running workers
		return 1
	}

	return utilizedRate
}

func (m *inputMetrics) pollSQSUtilizationMetric(ctx context.Context, maxMessagesInflight int) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	lastTick := time.Now()
	for {
		select {
		case <-ctx.Done():
			return
		case tick := <-t.C:
			duration := tick.Sub(lastTick)
			utilizedRate := calculateUtilizationAndReset(duration, maxMessagesInflight, m)
			m.sqsWorkerUtilization.Set(utilizedRate)
			// reset for the next polling duration
			lastTick = tick
		}
	}
}

func newInputMetrics(id string, optionalParent *monitoring.Registry, maxWorkers int) *inputMetrics {
	reg, unreg := inputmon.NewInputRegistry(inputName, id, optionalParent)
	ctx, cancelInputCtx := context.WithCancel(context.Background())

	out := &inputMetrics{
		registry:                            reg,
		unregister:                          unreg,
		ctx:                                 ctx,
		cancel:                              cancelInputCtx,
		sqsMessagesReceivedTotal:            monitoring.NewUint(reg, "sqs_messages_received_total"),
		sqsVisibilityTimeoutExtensionsTotal: monitoring.NewUint(reg, "sqs_visibility_timeout_extensions_total"),
		sqsMessagesInflight:                 monitoring.NewUint(reg, "sqs_messages_inflight_gauge"),
		sqsMessagesReturnedTotal:            monitoring.NewUint(reg, "sqs_messages_returned_total"),
		sqsMessagesDeletedTotal:             monitoring.NewUint(reg, "sqs_messages_deleted_total"),
		sqsWorkerUtilization:                monitoring.NewFloat(reg, "sqs_worker_utilization"),
		sqsMessageProcessingTime:            metrics.NewUniformSample(1024),
		sqsLagTime:                          metrics.NewUniformSample(1024),
		s3ObjectsRequestedTotal:             monitoring.NewUint(reg, "s3_objects_requested_total"),
		s3ObjectsAckedTotal:                 monitoring.NewUint(reg, "s3_objects_acked_total"),
		s3ObjectsListedTotal:                monitoring.NewUint(reg, "s3_objects_listed_total"),
		s3ObjectsProcessedTotal:             monitoring.NewUint(reg, "s3_objects_processed_total"),
		s3BytesProcessedTotal:               monitoring.NewUint(reg, "s3_bytes_processed_total"),
		s3EventsCreatedTotal:                monitoring.NewUint(reg, "s3_events_created_total"),
		s3ObjectsInflight:                   monitoring.NewUint(reg, "s3_objects_inflight_gauge"),
		s3ObjectProcessingTime:              metrics.NewUniformSample(1024),
	}
	adapter.NewGoMetrics(reg, "sqs_message_processing_time", adapter.Accept).
		Register("histogram", metrics.NewHistogram(out.sqsMessageProcessingTime)) //nolint:errcheck // A unique namespace is used so name collisions are impossible.
	adapter.NewGoMetrics(reg, "sqs_lag_time", adapter.Accept).
		Register("histogram", metrics.NewHistogram(out.sqsLagTime)) //nolint:errcheck // A unique namespace is used so name collisions are impossible.
	adapter.NewGoMetrics(reg, "s3_object_processing_time", adapter.Accept).
		Register("histogram", metrics.NewHistogram(out.s3ObjectProcessingTime)) //nolint:errcheck // A unique namespace is used so name collisions are impossible.

	go out.pollSQSUtilizationMetric(ctx, maxWorkers)

	return out
}

// monitoredReader implements io.Reader and counts the number of bytes read.
type monitoredReader struct {
	reader         io.Reader
	totalBytesRead *monitoring.Uint
}

func newMonitoredReader(r io.Reader, metric *monitoring.Uint) *monitoredReader {
	return &monitoredReader{reader: r, totalBytesRead: metric}
}

func (m *monitoredReader) Read(p []byte) (int, error) {
	n, err := m.reader.Read(p)
	m.totalBytesRead.Add(uint64(n))
	return n, err
}
