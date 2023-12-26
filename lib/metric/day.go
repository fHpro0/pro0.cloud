package metrics

import (
	"time"
)

type day struct {
	date           time.Time
	requests       int
	requestsNew    int
	requestsUpdate int
	requestsDelete int
	requestsFetch  int
}

// todayOfLast7Days returns today
func (m *Metrics) todayOfLast7Days() *day {
	tNow := time.Now()
	now := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), 0, 0, 0, 0, time.Local)

	// First clean up the list
	m.cleanupLast7Days()

	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, d := range m.last7Days {
		diff := now.Sub(d.date)
		diffDays := int(diff.Hours()) / 24
		if diffDays == 0 {
			// Today
			return d
		}
	}

	// Day does not exist in list so we create it
	d := &day{
		date: now,
	}
	// Append today to list
	m.last7Days = append(m.last7Days, d)
	return d
}

// cleanupLast7Days removes days older than 7
func (m *Metrics) cleanupLast7Days() {
	tNow := time.Now()
	now := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), 0, 0, 0, 0, time.Local)

	m.mutex.Lock()

	l := len(m.last7Days) - 1
	for i, d := range m.last7Days {
		diff := now.Sub(d.date)
		diffDays := int(diff.Hours()) / 24
		if diffDays >= 7 {
			// Delete days older than 7 days
			m.last7Days = append(m.last7Days[:i], m.last7Days[i+1:]...)
			m.mutex.Unlock()
			if i < l {
				m.cleanupLast7Days()
			}
			return
		}
	}
	// Unlock or wait for the end of the world
	m.mutex.Unlock()
}
