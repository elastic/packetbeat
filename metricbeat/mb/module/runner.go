package module

import (
	"sync"

	"github.com/elastic/beats/libbeat/publisher/beat"
)

// Runner is a facade for a Wrapper that provides a simple interface
// for starting and stopping a Module.
type Runner interface {
	// Start starts the Module. If Start is called more than once, only the
	// first will start the Module.
	Start()

	// Stop stops the Module and waits for module's MetricSets to exit. The
	// publisher.Client will be closed by Stop. If Stop is called more than
	// once, only the first stop the Module and wait for it to exit.
	Stop()

	// Added to be consistent with cfgfile.Runner
	ID() uint64
}

// NewRunner returns a Runner facade. The events generated by
// the Module will be published to a new publisher.Client generated from the
// pubClientFactory.
func NewRunner(client beat.Client, mod *Wrapper) Runner {
	return &runner{
		done:   make(chan struct{}),
		mod:    mod,
		client: client,
	}
}

type runner struct {
	done      chan struct{}
	wg        sync.WaitGroup
	startOnce sync.Once
	stopOnce  sync.Once
	mod       *Wrapper
	client    beat.Client
}

func (mr *runner) Start() {
	mr.startOnce.Do(func() {
		output := mr.mod.Start(mr.done)
		mr.wg.Add(1)
		go func() {
			defer mr.wg.Done()
			PublishChannels(mr.client, output)
		}()
	})
}

func (mr *runner) Stop() {
	mr.stopOnce.Do(func() {
		close(mr.done)
		mr.client.Close()
		mr.wg.Wait()
	})
}

func (mr *runner) ID() uint64 {
	return mr.mod.Hash()
}
