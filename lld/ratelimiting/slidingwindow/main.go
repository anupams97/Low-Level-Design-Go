package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rateLimiter := newSlidingWindow(3, time.Second)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 15; i++ {
		// Randomize the interval between 50ms and 200ms
		randomInterval := time.Duration(rand.Intn(150)+50) * time.Millisecond
		time.Sleep(randomInterval)

		if rateLimiter.IsAllowed() {
			fmt.Printf("Request %d success at %v\n", i+1, time.Now().Format("15:04:05.000"))
		} else {
			fmt.Printf("Request %d failed at %v\n", i+1, time.Now().Format("15:04:05.000"))
		}
	}
}

type slidingWindow struct {
	mu           sync.Mutex
	limit        int
	windowSize   time.Duration
	prevWinCount int
	curWinCount  int
	windowStart  time.Time
}

func (sw *slidingWindow) IsAllowed() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(sw.windowStart)

	if elapsed > sw.windowSize {

		overlap := sw.windowSize - elapsed%sw.windowSize
		overlapRatio := float64(overlap) / float64(sw.windowSize)

		sw.prevWinCount = int(float64(sw.curWinCount) * overlapRatio)
		sw.curWinCount = 0
		sw.windowStart = now.Truncate(sw.windowSize)
	}

	effectiveRate := sw.prevWinCount + sw.curWinCount
	if effectiveRate < sw.limit {
		sw.curWinCount++
		return true
	}

	return false
}

func newSlidingWindow(limit int, windowSize time.Duration) *slidingWindow {
	return &slidingWindow{
		limit:        limit,
		windowSize:   windowSize,
		prevWinCount: 0,
		windowStart:  time.Now(),
		curWinCount:  0,
	}
}
