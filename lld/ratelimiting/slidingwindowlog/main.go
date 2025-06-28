package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindowLog struct {
	limit      int
	windowSize time.Duration
	log        []time.Time
	mu         sync.Mutex
}

func NewSlidingWindowLog(limit int, windowSize time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		limit:      limit,
		windowSize: windowSize,
		log:        []time.Time{},
	}
}

func (sw *SlidingWindowLog) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-sw.windowSize)

	newLog := []time.Time{}
	for _, timestamp := range sw.log {
		if timestamp.After(windowStart) {
			newLog = append(newLog, timestamp)
		}
	}
	sw.log = newLog

	if len(sw.log) < sw.limit {
		sw.log = append(sw.log, now)
		return true
	}

	return false
}

func main() {
	limiter := NewSlidingWindowLog(5, 10*time.Second)

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %d allowed at %v\n", i+1, time.Now())
		} else {
			fmt.Printf("Request %d denied at %v\n", i+1, time.Now())
		}
		time.Sleep(2 * time.Second)
	}
}
