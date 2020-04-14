package counter

import "sync"

// Counter can be used to set, increment and get a thread-safe counter. Note:
// for learning purposes, this implementation is flawed with several bugs.
type Counter struct {
	sync.RWMutex
	counter int64
}

// New returns a counter initialized to zero (0)
func New() *Counter {
	return &Counter{}
}

// Set sets the counter to the specfified new value
func (c *Counter) Set(new int64) {
	c.Lock()
	defer c.Unlock()
	c.counter = new
}

// Get returns the current value of the counter
func (c *Counter) Get() int64 {
	c.RLock()
	defer c.Unlock()
	return c.counter
}

// Reset resets the counter to zero (0)
func (c *Counter) Reset() {
	c.Lock()
	defer c.Unlock()
	c.counter = 0
}

// Increment increments the counter by one
func (c *Counter) Increment() {
	current := c.Get()
	c.Set(current + 1)
}
