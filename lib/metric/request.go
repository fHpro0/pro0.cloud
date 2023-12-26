package metrics

import "fmt"

type RequestType uint8

const (
	RequestNew RequestType = iota
	DbError
)

func (m *Metrics) AddDb(t RequestType) {
	// Get actual hour
	h := m.actualOfLast6Hours()

	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Count up
	switch t {
	case DbError:
		h.dbError = h.dbError + 1
		break
	}
}

// DbLast6Hours returns the number of database errors of the last 6 hours.
func (m *Metrics) DbLast6Hours() (str string, err int) {
	// First clean up the list
	m.cleanupLast6Hours()

	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, day := range m.last6Hours {
		err = err + day.dbError
	}
	str = fmt.Sprintf("error: %d", err)

	return str, err
}
