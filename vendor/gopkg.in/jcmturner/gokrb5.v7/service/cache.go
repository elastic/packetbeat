// Package service provides server side integrations for Kerberos authentication.
package service

import (
	"gopkg.in/jcmturner/gokrb5.v7/types"
	"sync"
	"time"
)

/*The server MUST utilize a replay cache to remember any authenticator
presented within the allowable clock skew.
The replay cache will store at least the server name, along with the
client name, time, and microsecond fields from the recently-seen
authenticators, and if a matching tuple is found, the
KRB_AP_ERR_REPEAT error is returned.  Note that the rejection here is
restricted to authenticators from the same principal to the same
server.  Other client principals communicating with the same server
principal should not have their authenticators rejected if the time
and microsecond fields happen to match some other client's
authenticator.

If a server loses track of authenticators presented within the
allowable clock skew, it MUST reject all requests until the clock
skew interval has passed, providing assurance that any lost or
replayed authenticators will fall outside the allowable clock skew
and can no longer be successfully replayed.  If this were not done,
an attacker could subvert the authentication by recording the ticket
and authenticator sent over the network to a server and replaying
them following an event that caused the server to lose track of
recently seen authenticators.*/

// Cache for tickets received from clients keyed by fully qualified client name. Used to track replay of tickets.
type Cache struct {
	entries map[string]clientEntries
	mux     sync.RWMutex
}

// clientEntries holds entries of client details sent to the service.
type clientEntries struct {
	replayMap map[time.Time]replayCacheEntry
	seqNumber int64
	subKey    types.EncryptionKey
}

// Cache entry tracking client time values of tickets sent to the service.
type replayCacheEntry struct {
	presentedTime time.Time
	sName         types.PrincipalName
	cTime         time.Time // This combines the ticket's CTime and Cusec
}

func (c *Cache) getClientEntries(cname types.PrincipalName) (clientEntries, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	ce, ok := c.entries[cname.PrincipalNameString()]
	return ce, ok
}

func (c *Cache) getClientEntry(cname types.PrincipalName, t time.Time) (replayCacheEntry, bool) {
	if ce, ok := c.getClientEntries(cname); ok {
		c.mux.RLock()
		defer c.mux.RUnlock()
		if e, ok := ce.replayMap[t]; ok {
			return e, true
		}
	}
	return replayCacheEntry{}, false
}

// Instance of the ServiceCache. This needs to be a singleton.
var replayCache Cache
var once sync.Once

// GetReplayCache returns a pointer to the Cache singleton.
func GetReplayCache(d time.Duration) *Cache {
	// Create a singleton of the ReplayCache and start a background thread to regularly clean out old entries
	once.Do(func() {
		replayCache = Cache{
			entries: make(map[string]clientEntries),
		}
		go func() {
			for {
				// TODO consider using a context here.
				time.Sleep(d)
				replayCache.ClearOldEntries(d)
			}
		}()
	})
	return &replayCache
}

// AddEntry adds an entry to the Cache.
func (c *Cache) AddEntry(sname types.PrincipalName, a types.Authenticator) {
	ct := a.CTime.Add(time.Duration(a.Cusec) * time.Microsecond)
	if ce, ok := c.getClientEntries(a.CName); ok {
		c.mux.Lock()
		defer c.mux.Unlock()
		ce.replayMap[ct] = replayCacheEntry{
			presentedTime: time.Now().UTC(),
			sName:         sname,
			cTime:         ct,
		}
		ce.seqNumber = a.SeqNumber
		ce.subKey = a.SubKey
	} else {
		c.mux.Lock()
		defer c.mux.Unlock()
		c.entries[a.CName.PrincipalNameString()] = clientEntries{
			replayMap: map[time.Time]replayCacheEntry{
				ct: {
					presentedTime: time.Now().UTC(),
					sName:         sname,
					cTime:         ct,
				},
			},
			seqNumber: a.SeqNumber,
			subKey:    a.SubKey,
		}
	}
}

// ClearOldEntries clears entries from the Cache that are older than the duration provided.
func (c *Cache) ClearOldEntries(d time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for ke, ce := range c.entries {
		for k, e := range ce.replayMap {
			if time.Now().UTC().Sub(e.presentedTime) > d {
				delete(ce.replayMap, k)
			}
		}
		if len(ce.replayMap) == 0 {
			delete(c.entries, ke)
		}
	}
}

// IsReplay tests if the Authenticator provided is a replay within the duration defined. If this is not a replay add the entry to the cache for tracking.
func (c *Cache) IsReplay(sname types.PrincipalName, a types.Authenticator) bool {
	ct := a.CTime.Add(time.Duration(a.Cusec) * time.Microsecond)
	if e, ok := c.getClientEntry(a.CName, ct); ok {
		if e.sName.Equal(sname) {
			return true
		}
	}
	c.AddEntry(sname, a)
	return false
}
