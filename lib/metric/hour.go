package metrics

import (
	"time"
)

type hour struct {
	date             time.Time
	dbError          int
	authUnauthorized int
	authBadRequest   int
	authForbidden    int
	authError        int
}

// actualOfLast6Hours returns the actual hour
func (m *Metrics) actualOfLast6Hours() *hour {
	tNow := time.Now()
	now := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), tNow.Hour(), 0, 0, 0, time.Local)

	// First clean up the list
	m.cleanupLast6Hours()

	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, h := range m.last6Hours {
		diff := now.Sub(h.date)
		diffHours := int(diff.Hours())
		if diffHours == 0 {
			// Actual
			return h
		}
	}

	// Hour does not exist in list so we create it
	h := &hour{
		date: now,
	}
	// Append actual hour to list
	m.last6Hours = append(m.last6Hours, h)
	return h
}

// cleanupLast6Hours removes hours older than 6
func (m *Metrics) cleanupLast6Hours() {
	tNow := time.Now()
	now := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), tNow.Hour(), 0, 0, 0, time.Local)
	m.mutex.Lock()

	l := len(m.last6Hours) - 1
	for i, h := range m.last6Hours {
		h.date.Hour()
		diff := now.Sub(h.date)
		diffHours := int(diff.Hours())
		if diffHours >= 6 {
			// Delete hours older than 6 hours
			m.last6Hours = append(m.last6Hours[:i], m.last6Hours[i+1:]...)
			m.mutex.Unlock()
			if i < l {
				m.cleanupLast6Hours()
			}
			return
		}
	}
	// Unlock or wait for the end of the world
	m.mutex.Unlock()
}
