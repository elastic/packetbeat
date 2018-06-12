// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("heartbeat/fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
