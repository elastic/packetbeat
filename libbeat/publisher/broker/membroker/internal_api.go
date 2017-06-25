package membroker

import "github.com/elastic/beats/libbeat/publisher"

// producer -> broker API

type pushRequest struct {
	event publisher.Event
	seq   uint32
	state *produceState
}

type producerCancelRequest struct {
	state *produceState
	resp  chan producerCancelResponse
}

type producerCancelResponse struct {
	removed int
}

// consumer -> broker API

type getRequest struct {
	sz   int              // request sz events from the broker
	resp chan getResponse // channel to send response to
}

type getResponse struct {
	buf []publisher.Event
	ack *ackChan
}

type batchAckRequest struct{}

type batchCancelRequest struct{ ack *ackChan }
