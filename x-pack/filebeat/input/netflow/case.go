// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package netflow

import (
	"sync"
	"unicode"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/x-pack/filebeat/input/netflow/decoder/record"
)

var fieldNameConverter = caseConverter{
	conversion: make(map[string]string),
}

type caseConverter struct {
	rwMutex    sync.RWMutex
	conversion map[string]string
}

func (c *caseConverter) memoize(nfName, converted string) string {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.conversion[nfName] = converted
	return converted
}

func (c *caseConverter) ToSnakeCase(orig record.Map) common.MapStr {
	result := common.MapStr(make(map[string]interface{}, len(orig)))
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	for nfName, value := range orig {
		name, found := c.conversion[nfName]
		if !found {
			c.rwMutex.RUnlock()
			name = c.memoize(nfName, CamelCaseToSnakeCase(nfName))
			c.rwMutex.RLock()
		}
		result[name] = value
	}
	return result
}

// CamelCaseToSnakeCase converts a camel-case identifier to snake-case
// format. This function is tailored to some specifics of NetFlow field names.
// Don't reuse it.
func CamelCaseToSnakeCase(in string) string {
	out := make([]rune, 0, len(in)+4)
	runes := []rune(in)
	upperStrike := 1
	for pos, r := range runes {
		lr := unicode.ToLower(r)
		isUpper := lr != r
		if isUpper {
			if upperStrike == 0 {
				out = append(out, '_')
			}
			upperStrike++
		} else {
			if upperStrike > 2 {
				// Some magic here:
				// NetFlow usually lowercases all but the first letter of an
				// acronym (Icmp) Except when it is 2 characters long: (IP).
				// In other cases, it keeps all caps, but if we have a run of
				// more than 2 uppercase chars, then the last char belongs to
				// the next word:
				// postNATSourceIPv4Address     : post_nat_source_ipv4_address
				// selectorIDTotalFlowsObserved : selector_id_total_flows_...
				out = append(out, '_')
				out[pos], out[pos-1] = out[pos-1], out[pos]
			}
			upperStrike = 0
		}
		out = append(out, lr)
	}
	return string(out)
}
