package pool

import (
	s "github.com/elastic/beats/metricbeat/schema"
	c "github.com/elastic/beats/metricbeat/schema/mapstriface"
)

var (
	schema = s.Schema{
		"name": c.Str("pool"),
		"process_manager": c.Str("process manager"),
		"slow_requests": c.Int("slow requests"),
		"connections": s.Object{
			"accepted": c.Int("accepted conn"),
			"max_listen_queue": c.Int("max listen queue"),
			"max_queue_limit": c.Int("listen queue len"),
			"queued": c.Int("listen queue"),
		},
		"processes": s.Object{
			"active": c.Int("active processes"),
			"idle":   c.Int("idle processes"),
			"max_active": c.Int("max active processes"),
			"max_children_reached": c.Int("max children reached"),
			"total": c.Int("total processes"),
		},
		"start_time": c.Str("start time"),
		"start_since": c.Int("start since"),
	}
)
