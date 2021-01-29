// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package browser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/x-pack/heartbeat/monitors/browser/source"
)

func TestValidInline(t *testing.T) {
	script := "a script"
	testParams := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	cfg := common.MustNewConfigFrom(common.MapStr{
		"name":   "My Name",
		"id":     "myId",
		"params": testParams,
		"source": common.MapStr{
			"inline": common.MapStr{
				"script": script,
			},
		},
	})
	s, e := NewSuite(cfg)
	require.NoError(t, e)
	require.NotNil(t, s)
	sSrc, ok := s.InlineSource()
	require.True(t, ok)
	require.Equal(t, script, script, sSrc)
	require.Equal(t, "", s.Workdir())
	require.Equal(t, testParams, s.Params())
}

func TestNameRequired(t *testing.T) {
	cfg := common.MustNewConfigFrom(common.MapStr{
		"id": "myId",
		"source": common.MapStr{
			"inline": common.MapStr{
				"script": "a script",
			},
		},
	})
	_, e := NewSuite(cfg)
	require.Regexp(t, ErrNameRequired, e)
}

func TestIDRequired(t *testing.T) {
	cfg := common.MustNewConfigFrom(common.MapStr{
		"name": "My Name",
		"source": common.MapStr{
			"inline": common.MapStr{
				"script": "a script",
			},
		},
	})
	_, e := NewSuite(cfg)
	require.Regexp(t, ErrIdRequired, e)
}

func TestEmptySource(t *testing.T) {
	cfg := common.MustNewConfigFrom(common.MapStr{
		"source": common.MapStr{},
	})
	s, e := NewSuite(cfg)

	require.Regexp(t, ErrBadConfig(source.ErrInvalidSource), e)
	require.Nil(t, s)
}
