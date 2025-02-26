package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value uint64
}

func (c *Counter) Add(a uint64) {
	c.mu.Lock()
	c.value += a
	c.mu.Unlock()
}

func (c *Counter) Get() uint64 {
	return c.value
}

func (c *Counter) GetAndReset() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	value := c.value
	c.value = 0

	return value
}

type Counters struct {
	Datagrams *Counter
}

func NewCounters() *Counters {
	return &Counters{
		Datagrams: &Counter{},
	}
}

func (c *Counters) GetStringAndReset() string {
	return fmt.Sprintf("Datagrams forwarded: %v", c.Datagrams.GetAndReset())
}
