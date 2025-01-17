// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package es

import (
	"sync"

	conf "github.com/elastic/elastic-agent-libs/config"
)

type OnConfigUpdateFunc func(c *conf.C)
type UnsubscribeFunc func()

type Notifier struct {
	mx sync.RWMutex

	listeners map[int]OnConfigUpdateFunc
	id        int
}

func NewNotifier() *Notifier {
	return &Notifier{
		listeners: make(map[int]OnConfigUpdateFunc),
		id:        0,
	}
}

func (n *Notifier) Subscribe(fn OnConfigUpdateFunc) UnsubscribeFunc {
	n.mx.Lock()
	defer n.mx.Unlock()

	id := n.id
	n.id++
	n.listeners[id] = fn

	return func() {
		n.mx.Lock()
		defer n.mx.Unlock()
		delete(n.listeners, id)
	}
}

func (n *Notifier) Notify(c *conf.C) {
	n.mx.RLock()
	defer n.mx.RUnlock()

	for _, listener := range n.listeners {
		go listener(c)
	}
}
