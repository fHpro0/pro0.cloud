package metrics

import (
	"sync"
	"time"
)

type Metrics struct {
	timeServerStart time.Time
	last7Days       []*day
	last6Hours      []*hour
	mutex           sync.Mutex
}

// NewMetrics returns a new Export instance.
func NewMetrics() *Metrics {
	m := &Metrics{
		timeServerStart: time.Now(),
	}

	return m
}
