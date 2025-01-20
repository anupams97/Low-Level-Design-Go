package main

import (
	"fmt"
	"time"
)

func main() {
	rateLimiter := NewFixedWindow(3, time.Second*1)
	for i := 0; i < 15; i++ {
		time.Sleep(time.Millisecond * 100)
		if rateLimiter.IsAllowed() {
			fmt.Printf("Request %d success at %v\n", i+1, time.Now().Second())
		} else {
			fmt.Printf("Request %d failed at %v\n", i+1, time.Now().Second())
		}
	}
}

type FixedWindow struct {
	WindowStart time.Time
	WindowSize  time.Duration
	Limit       int
	Requests    int
}

func (fw *FixedWindow) IsAllowed() bool {
	now := time.Now()
	if now.Sub(fw.WindowStart) > fw.WindowSize {
		fw.WindowStart = time.Now()
		fw.Requests = 0
	}
	if fw.Requests < fw.Limit {
		fw.Requests++
		return true
	}
	return false
}

func NewFixedWindow(limit int, windowSize time.Duration) *FixedWindow {
	return &FixedWindow{
		Limit:       limit,
		WindowSize:  windowSize,
		Requests:    0,
		WindowStart: time.Now(),
	}
}
