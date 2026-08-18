package metrics

import "sync"

type Counter struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewCounter() *Counter {
	return &Counter{values: make(map[string]int)}
}

func (c *Counter) Add(name string, amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[name] += amount
}

func (c *Counter) Get(name string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.values[name]
}

func (c *Counter) Snapshot() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	copyOfValues := make(map[string]int, len(c.values))
	for name, value := range c.values {
		copyOfValues[name] = value
	}
	return copyOfValues
}
