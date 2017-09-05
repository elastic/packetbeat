package memqueue

import (
	"sync"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/publisher/queue"
)

type Broker struct {
	done chan struct{}

	logger logger

	bufSize int
	// buf         brokerBuffer
	// minEvents   int
	// idleTimeout time.Duration

	// api channels
	events    chan pushRequest
	requests  chan getRequest
	pubCancel chan producerCancelRequest

	// internal channels
	acks          chan int
	scheduledACKs chan chanList

	eventer queue.Eventer

	// wait group for worker shutdown
	wg          sync.WaitGroup
	waitOnClose bool
}

type Settings struct {
	Eventer        queue.Eventer
	Events         int
	FlushMinEvents int
	FlushTimeout   time.Duration
	WaitOnClose    bool
}

type ackChan struct {
	next         *ackChan
	ch           chan batchAckMsg
	seq          uint
	start, count int // number of events waiting for ACK
	states       []clientState
}

type chanList struct {
	head *ackChan
	tail *ackChan
}

func init() {
	queue.RegisterType("mem", create)
}

func create(eventer queue.Eventer, cfg *common.Config) (queue.Queue, error) {
	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, err
	}

	return NewBroker(Settings{
		Eventer:        eventer,
		Events:         config.Events,
		FlushMinEvents: config.FlushMinEvents,
		FlushTimeout:   config.FlushTimeout,
	}), nil
}

// NewBroker creates a new broker based in-memory queue holding up to sz number of events.
// If waitOnClose is set to true, the broker will block on Close, until all internal
// workers handling incoming messages and ACKs have been shut down.
func NewBroker(
	settings Settings,
) *Broker {
	// define internal channel size for procuder/client requests
	// to the broker
	chanSize := 20

	var (
		sz           = settings.Events
		minEvents    = settings.FlushMinEvents
		flushTimeout = settings.FlushTimeout
	)

	if minEvents < 1 {
		minEvents = 1
	}
	if minEvents > 1 && flushTimeout <= 0 {
		minEvents = 1
		flushTimeout = 0
	}
	if minEvents > sz {
		minEvents = sz
	}

	logger := defaultLogger
	b := &Broker{
		done:   make(chan struct{}),
		logger: logger,

		// broker API channels
		events:    make(chan pushRequest, chanSize),
		requests:  make(chan getRequest),
		pubCancel: make(chan producerCancelRequest, 5),

		// internal broker and ACK handler channels
		acks:          make(chan int),
		scheduledACKs: make(chan chanList),

		waitOnClose: settings.WaitOnClose,

		eventer: settings.Eventer,
	}

	eventLoop := newEventLoop(b, sz, minEvents, flushTimeout)
	b.bufSize = sz
	ack := &ackLoop{broker: b}

	b.wg.Add(2)
	go func() {
		defer b.wg.Done()
		eventLoop.run()
	}()
	go func() {
		defer b.wg.Done()
		ack.run()
	}()

	return b
}

func (b *Broker) Close() error {
	close(b.done)
	if b.waitOnClose {
		b.wg.Wait()
	}
	return nil
}

func (b *Broker) BufferConfig() queue.BufferConfig {
	return queue.BufferConfig{
		Events: b.bufSize,
	}
}

func (b *Broker) Producer(cfg queue.ProducerConfig) queue.Producer {
	return newProducer(b, cfg.ACK, cfg.OnDrop, cfg.DropOnCancel)
}

func (b *Broker) Consumer() queue.Consumer {
	return newConsumer(b)
}

var ackChanPool = sync.Pool{
	New: func() interface{} {
		return &ackChan{
			ch: make(chan batchAckMsg, 1),
		}
	},
}

func newACKChan(seq uint, start, count int, states []clientState) *ackChan {
	ch := ackChanPool.Get().(*ackChan)
	ch.next = nil
	ch.seq = seq
	ch.start = start
	ch.count = count
	ch.states = states
	return ch
}

func releaseACKChan(c *ackChan) {
	c.next = nil
	ackChanPool.Put(c)
}

func (l *chanList) prepend(ch *ackChan) {
	ch.next = l.head
	l.head = ch
	if l.tail == nil {
		l.tail = ch
	}
}

func (l *chanList) concat(other *chanList) {
	if l.head == nil {
		*l = *other
		return
	}

	l.tail.next = other.head
	l.tail = other.tail
}

func (l *chanList) append(ch *ackChan) {
	if l.head == nil {
		l.head = ch
	} else {
		l.tail.next = ch
	}
	l.tail = ch
}

func (l *chanList) count() (elems, count int) {
	for ch := l.head; ch != nil; ch = ch.next {
		elems++
		count += ch.count
	}
	return
}

func (l *chanList) empty() bool {
	return l.head == nil
}

func (l *chanList) front() *ackChan {
	return l.head
}

func (l *chanList) channel() chan batchAckMsg {
	if l.head == nil {
		return nil
	}
	return l.head.ch
}

func (l *chanList) pop() *ackChan {
	ch := l.head
	if ch != nil {
		l.head = ch.next
		if l.head == nil {
			l.tail = nil
		}
	}

	ch.next = nil
	return ch
}
