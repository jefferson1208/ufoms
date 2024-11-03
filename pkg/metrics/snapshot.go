package metrics

import "time"

type Metadata map[string]any

type Snapshot struct {
	ID            string
	CorrelationID string
	Component     string
	Start         time.Time
	End           time.Time
	Duration      time.Duration
	Metadata      Metadata
	Parent        *Snapshot
}
