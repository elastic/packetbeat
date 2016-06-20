// +build darwin linux openbsd windows

package memory

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
	system "github.com/elastic/beats/metricbeat/module/system/common"

	"github.com/pkg/errors"
)

func init() {
	if err := mb.Registry.AddMetricSet("system", "memory", New); err != nil {
		panic(err)
	}
}

// MetricSet for fetching system memory metrics.
type MetricSet struct {
	mb.BaseMetricSet
}

// New is a mb.MetricSetFactory that returns a memory.MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	return &MetricSet{base}, nil
}

// Fetch fetches memory metrics from the OS.
func (m *MetricSet) Fetch() (event common.MapStr, err error) {
	memStat, err := system.GetMemory()
	if err != nil {
		return nil, errors.Wrap(err, "memory")
	}
	system.AddMemPercentage(memStat)

	swapStat, err := system.GetSwap()
	if err != nil {
		return nil, errors.Wrap(err, "swap")
	}
	system.AddSwapPercentage(swapStat)

	memory := common.MapStr{
		"total": memStat.Total,
		"used": common.MapStr{
			"bytes": memStat.Used,
			"pct":   memStat.UsedPercent,
		},
		"free": memStat.Free,
		"actual": common.MapStr{
			"free": memStat.ActualFree,
			"used": common.MapStr{
				"pct":   memStat.ActualUsedPercent,
				"bytes": memStat.ActualUsed,
			},
		},
	}

	swap := common.MapStr{
		"total": swapStat.Total,
		"used": common.MapStr{
			"bytes": swapStat.Used,
			"pct":   swapStat.UsedPercent,
		},
		"free": swapStat.Free,
	}

	memory["swap"] = swap
	return memory, nil
}
