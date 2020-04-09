package counter

import "sync"

type counter struct {
	count int
	//A Mutex is a mutual exclusion lock.
	//The zero value for a Mutex is an unlocked mutex.
	mu sync.Mutex
}

func Counter() *counter {
	return &counter{}
}

func (c *counter) Value() int {
	return c.count
}

func (c *counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
