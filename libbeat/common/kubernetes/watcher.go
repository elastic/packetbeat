package kubernetes

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/ericchiang/k8s"
	"github.com/ericchiang/k8s/apis/core/v1"

	"github.com/elastic/beats/libbeat/logp"
)

func filterByNode(node string) k8s.Option {
	return k8s.QueryParam("fieldSelector", "spec.nodeName="+node)
}

// Watcher watches Kubernetes resources events
type Watcher interface {
	// Start watching Kubernetes API for new events after resources were listed
	Start() error

	// Stop watching Kubernetes API for new events
	Stop()

	// AddEventHandler add event handlers for corresponding event type watched
	AddEventHandler(ResourceEventHandler)
}

// WatchOptions controls watch behaviors
type WatchOptions struct {
	// SyncTimeout is a timeout for listing historical resources
	SyncTimeout time.Duration
	// Node is used for filtering watched resource to given node, use "" for all nodes
	Node string
	// Namespace is used for filtering watched resource to given namespace, use "" for all namespaces
	Namespace string
}

type watcher struct {
	client              *k8s.Client
	options             WatchOptions
	lastResourceVersion string
	ctx                 context.Context
	stop                context.CancelFunc
	resourceList        k8s.ResourceList
	k8sResourceFactory  func() k8s.Resource
	resourceFactory     func() Resource
	handler             ResourceEventHandler
}

// NewWatcher initializes the watcher client to provide a events handler for
// resource from the cluster (filtered to the given node)
func NewWatcher(client *k8s.Client, resource Resource, options WatchOptions) (Watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{
		client:              client,
		options:             options,
		lastResourceVersion: "0",
		ctx:                 ctx,
		stop:                cancel,
	}
	switch resource.(type) {
	case *Pod:
		w.resourceList = &v1.PodList{}
		w.k8sResourceFactory = func() k8s.Resource { return &v1.Pod{} }
		w.resourceFactory = func() Resource { return &Pod{} }
	case *Event:
		w.resourceList = &v1.EventList{}
		w.k8sResourceFactory = func() k8s.Resource { return &v1.Event{} }
		w.resourceFactory = func() Resource { return &Event{} }
	default:
		return nil, fmt.Errorf("unsupported resource type for watching %T", resource)
	}
	return w, nil
}

func (w *watcher) AddEventHandler(h ResourceEventHandler) {
	w.handler = h
}

func (w *watcher) buildOpts() []k8s.Option {
	options := []k8s.Option{k8s.ResourceVersion(w.lastResourceVersion)}
	if w.options.Node != "" {
		options = append(options, filterByNode(w.options.Node))
	}
	return options
}

func (w *watcher) sync() error {
	ctx, cancel := context.WithTimeout(w.ctx, w.options.SyncTimeout)
	defer cancel()

	logp.Info("kubernetes: Performing a resource sync for %T", w.resourceList)

	err := w.client.List(ctx, w.options.Namespace, w.resourceList, w.buildOpts()...)
	if err != nil {
		logp.Err("kubernetes: Performing a resource sync err %s for %T", err.Error(), w.resourceList)
		return err
	}

	// ugly reflect since ResourceList does not exposed Items
	// can be avoid by using k8s.io/client-go
	list := reflect.ValueOf(w.resourceList).Elem().FieldByName("Items")
	for i := 0; i < list.Len(); i++ {
		w.onAdd(list.Index(i).Interface().(k8s.Resource))
	}

	// Store last version
	w.lastResourceVersion = w.resourceList.GetMetadata().GetResourceVersion()

	logp.Info("kubernetes: %s", "Resource sync done")
	return nil
}

func (w *watcher) onAdd(obj k8s.Resource) {
	w.handler.OnAdd(resourceConverter(obj, w.resourceFactory()))
}

func (w *watcher) onUpdate(obj k8s.Resource) {
	w.handler.OnUpdate(resourceConverter(obj, w.resourceFactory()))
}

func (w *watcher) onDelete(obj k8s.Resource) {
	w.handler.OnDelete(resourceConverter(obj, w.resourceFactory()))
}

// Start watching pods
func (w *watcher) Start() error {

	// Make sure that events don't flow into the annotator before informer is fully set up
	// Sync initial state:
	err := w.sync()
	if err != nil {
		w.Stop()
		return err
	}

	// Watch for new changes
	go w.watch()

	return nil
}

func (w *watcher) watch() {
	for {
		select {
		case <-w.ctx.Done():
			logp.Info("kubernetes: %s", "Watching API for resource events stopped")
			return
		default:
		}

		logp.Info("kubernetes: %s", "Watching API for resource events")

		watcher, err := w.client.Watch(w.ctx, w.options.Namespace, w.k8sResourceFactory(), w.buildOpts()...)
		if err != nil {
			//watch failures should be logged and gracefully failed over as metadata retrieval
			//should never stop.
			logp.Err("kubernetes: Watching API error %v", err)
			time.Sleep(time.Second)
			continue
		}

		for {
			r := w.k8sResourceFactory()
			eventType, err := watcher.Next(r)
			if err != nil {
				logp.Err("kubernetes: Watching API error %v", err)
				watcher.Close()
				break
			}
			switch eventType {
			case k8s.EventAdded:
				w.onAdd(r)
			case k8s.EventModified:
				w.onUpdate(r)
			case k8s.EventDeleted:
				w.onDelete(r)
			default:
				logp.Err("kubernetes: Watching API error with event type %s", eventType)
			}
		}
	}
}

func (w *watcher) Stop() {
	w.stop()
}
