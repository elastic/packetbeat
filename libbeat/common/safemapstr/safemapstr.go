package safemapstr

import (
	"strings"

	"github.com/elastic/beats/libbeat/common"
)

const alternativeKey = "value"

// Put This method implements a way to put dotted keys into a MapStr while
// ensuring they don't override each other. For example:
//
//  a := MapStr{}
//  safemapstr.Put(a, "com.docker.swarm.task", "x")
//  safemapstr.Put(a, "com.docker.swarm.task.id", 1)
//  safemapstr.Put(a, "com.docker.swarm.task.name", "foobar")
//
// Will result in `{"com":{"docker":{"swarm":{"task":{"id":1,"name":"foobar","value":"x"}}}}}`
//
// Put detects this scenario and renames the common base key, by appending
// `.value`
func Put(data common.MapStr, key string, value interface{}) error {
	// XXX This implementation mimics `common.MapStr.Put`, both should be updated to have similar behavior

	d, k := mapFind(data, key)
	d[k] = value
	return nil
}

// mapFind walk the map based on the given dotted key and returns the final map
// and key to operate on. This function adds intermediate maps, if the key is
// missing from the original map.

func mapFind(data common.MapStr, key string) (common.MapStr, string) {
	// XXX This implementation mimics `common.mapFind`, both should be updated to have similar behavior

	for {
		if oldValue, exists := data[key]; exists {
			if oldMap, ok := tryToMapStr(oldValue); ok {
				return oldMap, alternativeKey
			}
			return data, key
		}

		idx := strings.IndexRune(key, '.')
		if idx < 0 {
			// if old value exists and is a dictionary, return the old dictionary and
			// make sure we store the new value using the 'alternativeKey'
			if oldValue, exists := data[key]; exists {
				if oldMap, ok := tryToMapStr(oldValue); ok {
					return oldMap, alternativeKey
				}
			}

			return data, key
		}

		// Check if first sub-key exists. Create an intermediate map if not.
		k := key[:idx]
		d, exists := data[k]
		if !exists {
			d = common.MapStr{}
			data[k] = d
		}

		// store old value under 'alternativeKey' if the old value is no map.
		// Do not overwrite old value.
		v, ok := tryToMapStr(d)
		if !ok {
			v = common.MapStr{alternativeKey: d}
			data[k] = v
		}

		// advance into sub-map
		key = key[idx+1:]
		data = v
	}
}

func tryToMapStr(v interface{}) (common.MapStr, bool) {
	switch m := v.(type) {
	case common.MapStr:
		return m, true
	case map[string]interface{}:
		return common.MapStr(m), true
	default:
		return nil, false
	}
}
