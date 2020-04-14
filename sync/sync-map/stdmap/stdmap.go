package stdmap

import "sync"

// Map is a lock-protected map using a builtin map[string]string
type Map struct {
	sync.RWMutex
	internal map[int]int
}

// New creates a new Map
func New() *Map {
	return &Map{
		internal: make(map[int]int),
	}
}

// Load loads the value for the given key (if it exists) from the map
func (m *Map) Load(key int) (int, bool) {
	m.RLock()
	result, ok := m.internal[key]
	m.RUnlock()
	return result, ok
}

// Store stores the given key and value in the map
func (m *Map) Store(key, value int) {
	m.Lock()
	m.internal[key] = value
	m.Unlock()
}

// Delete deletes the given key from the map
func (m *Map) Delete(key int) {
	m.Lock()
	delete(m.internal, key)
	m.Unlock()
}
