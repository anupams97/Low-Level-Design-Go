package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int
	m     sync.Mutex
}

func (c *Counter) Increment() {
	defer c.m.Unlock()
	c.m.Lock()
	c.count++
}

func (c *Counter) Decrement() {
	defer c.m.Unlock()
	c.m.Lock()
	c.count--
}

func (c *Counter) GetCount() int {
	defer c.m.Unlock()
	c.m.Lock()
	return c.count
}

func (c *Counter) GetCountString() string {
	defer c.m.Unlock()
	c.m.Lock()
	return fmt.Sprintf("%d", c.count)
}

type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *AtomicCounter) Decrement() {
	atomic.AddInt64(&c.count, -1)
}

func (c *AtomicCounter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

func (c *AtomicCounter) GetCountString() string {
	return fmt.Sprintf("%d", atomic.LoadInt64(&c.count))
}
