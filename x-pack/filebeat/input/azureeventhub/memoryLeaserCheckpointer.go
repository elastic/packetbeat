package azureeventhub

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Azure/azure-event-hubs-go/v2"
	"github.com/Azure/azure-event-hubs-go/v2/eph"

	"github.com/Azure/azure-amqp-common-go/v2/uuid"
	"github.com/devigned/tab"

	"github.com/Azure/azure-event-hubs-go/v2/persist"
)

type (
	memoryLeaserCheckpointer struct {
		store         *sharedStore
		processor     *eph.EventProcessorHost
		leaseDuration time.Duration
		memMu         sync.Mutex
		leases        map[string]*memoryLease
	}

	memoryLease struct {
		eph.Lease
		expirationTime time.Time
		Token          string
		Checkpoint     *persist.Checkpoint
		leaser         *memoryLeaserCheckpointer
	}

	sharedStore struct {
		leases  map[string]*storeLease
		storeMu sync.Mutex
	}

	storeLease struct {
		token      string
		expiration time.Time
		ml         *memoryLease
	}
)

func newMemoryLease(partitionID string) *memoryLease {
	lease := new(memoryLease)
	lease.PartitionID = partitionID
	return lease
}

func (s *sharedStore) exists() bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	return s.leases != nil
}

func (s *sharedStore) ensure() bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if s.leases == nil {
		s.leases = make(map[string]*storeLease)
	}
	return true
}

func (s *sharedStore) getLease(partitionID string) memoryLease {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	return *s.leases[partitionID].ml
}

func (s *sharedStore) deleteLease(partitionID string) {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	delete(s.leases, partitionID)
}

func (s *sharedStore) createOrGetLease(partitionID string) memoryLease {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if _, ok := s.leases[partitionID]; !ok {
		s.leases[partitionID] = new(storeLease)
	}

	l := s.leases[partitionID]
	if l.ml != nil {
		return *l.ml
	}
	l.ml = newMemoryLease(partitionID)
	return *l.ml
}

func (s *sharedStore) changeLease(partitionID, newToken, oldToken string, duration time.Duration) bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if l, ok := s.leases[partitionID]; ok && l.token == oldToken {
		l.token = newToken
		l.expiration = time.Now().Add(duration)
		return true
	}
	return false
}

func (s *sharedStore) releaseLease(partitionID, token string) bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if l, ok := s.leases[partitionID]; ok && l.token == token {
		l.token = ""
		l.expiration = time.Now().Add(-1 * time.Second)
		return true
	}
	return false
}

func (s *sharedStore) renewLease(partitionID, token string, duration time.Duration) bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if l, ok := s.leases[partitionID]; ok && l.token == token {
		l.expiration = time.Now().Add(duration)
		return true
	}
	return false
}

func (s *sharedStore) acquireLease(partitionID, newToken string, duration time.Duration) bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if l, ok := s.leases[partitionID]; ok && (time.Now().After(l.expiration) || l.token == "") {
		l.token = newToken
		l.expiration = time.Now().Add(duration)
		return true
	}
	return false
}

func (s *sharedStore) storeLease(partitionID, token string, ml memoryLease) bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if l, ok := s.leases[partitionID]; ok && l.token == token {
		l.ml = &ml
		return true
	}
	return false
}

func (s *sharedStore) isLeased(partitionID string) bool {
	s.storeMu.Lock()
	defer s.storeMu.Unlock()

	if l, ok := s.leases[partitionID]; ok {
		if time.Now().After(l.expiration) || l.token == "" {
			return false
		}
		return true
	}
	return false
}

// IsNotOwnedOrExpired indicates that the lease has expired and does not owned by a processor
func (l *memoryLease) isNotOwnedOrExpired(ctx context.Context) bool {
	return l.IsExpired(ctx) || l.Owner == ""
}

// IsExpired indicates that the lease has expired and is no longer valid
func (l *memoryLease) IsExpired(_ context.Context) bool {
	return !l.leaser.store.isLeased(l.PartitionID)
}

func (l *memoryLease) expireAfter(d time.Duration) {
	l.expirationTime = time.Now().Add(d)
}

// GetEpoch returns the value of the epoch
func (l *memoryLease) GetEpoch() int64 {
	return l.Epoch
}

// NewMemoryLeaserCheckpointer will be used for testing, creates a new memory leaser/checkpointer
func NewMemoryLeaserCheckpointer(leaseDuration time.Duration, store *sharedStore) *memoryLeaserCheckpointer {
	return &memoryLeaserCheckpointer{
		leaseDuration: leaseDuration,
		leases:        make(map[string]*memoryLease),
		store:         store,
	}
}

// SetEventHostProcessor sets the event host processor
func (ml *memoryLeaserCheckpointer) SetEventHostProcessor(eph *eph.EventProcessorHost) {
	ml.processor = eph
}

// StoreExists checks if store exists
func (ml *memoryLeaserCheckpointer) StoreExists(ctx context.Context) (bool, error) {
	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.StoreExists")
	defer span.End()

	return ml.store.exists(), nil
}

// EnsureStore creates the store
func (ml *memoryLeaserCheckpointer) EnsureStore(ctx context.Context) error {
	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.EnsureStore")
	defer span.End()

	ml.store.ensure()
	return nil
}

// DeleteStore deletes the store
func (ml *memoryLeaserCheckpointer) DeleteStore(ctx context.Context) error {
	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.DeleteStore")
	defer span.End()

	return ml.EnsureStore(ctx)
}

// GetLeases gets the lease
func (ml *memoryLeaserCheckpointer) GetLeases(ctx context.Context) ([]eph.LeaseMarker, error) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.GetLeases")
	defer span.End()

	partitionIDs := ml.processor.GetPartitionIDs()
	leases := make([]eph.LeaseMarker, len(partitionIDs))
	for idx, partitionID := range partitionIDs {
		lease := ml.store.getLease(partitionID)
		lease.leaser = ml
		leases[idx] = &lease
	}
	return leases, nil
}

// EnsureLease creates or gets the lease
func (ml *memoryLeaserCheckpointer) EnsureLease(ctx context.Context, partitionID string) (eph.LeaseMarker, error) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.EnsureLease")
	defer span.End()

	l := ml.store.createOrGetLease(partitionID)
	l.leaser = ml
	return &l, nil
}

// DeleteLease deletes the lease
func (ml *memoryLeaserCheckpointer) DeleteLease(ctx context.Context, partitionID string) error {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.DeleteLease")
	defer span.End()

	ml.store.deleteLease(partitionID)
	return nil
}

// AcquireLease acquires the lease
func (ml *memoryLeaserCheckpointer) AcquireLease(ctx context.Context, partitionID string) (eph.LeaseMarker, bool, error) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, ctx := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.AcquireLease")
	defer span.End()

	lease := ml.store.getLease(partitionID)
	lease.leaser = ml
	uuidToken, err := uuid.NewV4()
	if err != nil {
		tab.For(ctx).Error(err)
		return nil, false, err
	}

	newToken := uuidToken.String()
	if ml.store.isLeased(partitionID) {
		// is leased by someone else due to a race to acquire
		if !ml.store.changeLease(partitionID, newToken, lease.Token, ml.leaseDuration) {
			return nil, false, errors.New("failed to change lease")
		}
	} else {
		if !ml.store.acquireLease(partitionID, newToken, ml.leaseDuration) {
			return nil, false, errors.New("failed to acquire lease")
		}
	}

	lease.Token = newToken
	lease.Owner = ml.processor.GetName()
	lease.IncrementEpoch()
	if !ml.store.storeLease(partitionID, newToken, lease) {
		return nil, false, errors.New("failed to store lease after acquiring or changing")
	}
	ml.leases[partitionID] = &lease
	return &lease, true, nil
}

// RenewLease renews the lease
func (ml *memoryLeaserCheckpointer) RenewLease(ctx context.Context, partitionID string) (eph.LeaseMarker, bool, error) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.RenewLease")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if !ok {
		return nil, false, errors.New("lease was not found")
	}

	if !ml.store.renewLease(partitionID, lease.Token, ml.leaseDuration) {
		return nil, false, errors.New("unable to renew lease")
	}
	return lease, true, nil
}

// ReleaseLease releases the lease
func (ml *memoryLeaserCheckpointer) ReleaseLease(ctx context.Context, partitionID string) (bool, error) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.ReleaseLease")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if !ok {
		return false, errors.New("lease was not found")
	}

	if !ml.store.releaseLease(partitionID, lease.Token) {
		return false, errors.New("could not release the lease")
	}
	delete(ml.leases, partitionID)
	return true, nil
}

// UpdateLease updates the lease
func (ml *memoryLeaserCheckpointer) UpdateLease(ctx context.Context, partitionID string) (eph.LeaseMarker, bool, error) {
	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryLeaserCheckpointer.UpdateLease")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if !ok {
		return nil, false, errors.New("lease was not found")
	}

	if !ml.store.renewLease(partitionID, lease.Token, ml.leaseDuration) {
		return nil, false, errors.New("unable to renew lease")
	}

	if !ml.store.storeLease(partitionID, lease.Token, *lease) {
		return nil, false, errors.New("unable to store lease after renewal")
	}

	return lease, true, nil
}

// GetCheckpoint gets the checkpoint
func (ml *memoryLeaserCheckpointer) GetCheckpoint(ctx context.Context, partitionID string) (persist.Checkpoint, bool) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryCheckpointer.GetCheckpoint")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if ok {
		return *lease.Checkpoint, ok
	}
	return persist.NewCheckpointFromStartOfStream(), ok
}

// EnsureCheckpoint creates the checkpoint
func (ml *memoryLeaserCheckpointer) EnsureCheckpoint(ctx context.Context, partitionID string) (persist.Checkpoint, error) {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryCheckpointer.EnsureCheckpoint")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if ok {
		if lease.Checkpoint == nil {
			checkpoint := persist.NewCheckpointFromStartOfStream()
			lease.Checkpoint = &checkpoint
		}
		return *lease.Checkpoint, nil
	}
	return persist.NewCheckpointFromStartOfStream(), nil
}

// UpdateCheckpoint updates the checkpoint
func (ml *memoryLeaserCheckpointer) UpdateCheckpoint(ctx context.Context, partitionID string, checkpoint persist.Checkpoint) error {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryCheckpointer.UpdateCheckpoint")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if !ok {
		return errors.New("lease for partition isn't owned by this EventProcessorHost")
	}

	lease.Checkpoint = &checkpoint
	if !ml.store.storeLease(partitionID, lease.Token, *lease) {
		return errors.New("could not store lease on update of checkpoint")
	}
	return nil
}

// DeleteCheckpoint removes the checkpoint
func (ml *memoryLeaserCheckpointer) DeleteCheckpoint(ctx context.Context, partitionID string) error {
	ml.memMu.Lock()
	defer ml.memMu.Unlock()

	span, _ := startConsumerSpanFromContext(ctx, "eph.memoryCheckpointer.DeleteCheckpoint")
	defer span.End()

	lease, ok := ml.leases[partitionID]
	if !ok {
		return errors.New("lease for partition isn't owned by this EventProcessorHost")
	}

	checkpoint := persist.NewCheckpointFromStartOfStream()
	lease.Checkpoint = &checkpoint
	if !ml.store.storeLease(partitionID, lease.Token, *lease) {
		return errors.New("failed to store deleted checkpoint")
	}
	ml.leases[partitionID] = lease
	return nil
}

func (ml *memoryLeaserCheckpointer) Close() error {
	return nil
}

func startConsumerSpanFromContext(ctx context.Context, operationName string) (tab.Spanner, context.Context) {
	ctx, span := tab.StartSpan(ctx, operationName)
	eventhub.ApplyComponentInfo(span)
	span.AddAttributes(tab.StringAttribute("span.kind", "consumer"))
	return span, ctx
}
