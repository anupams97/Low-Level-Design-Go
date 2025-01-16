package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Todo rate limit/ memoization / logging
func main() {
	rand.Seed(time.Now().UnixNano())
	doWork := WithLogging(WithRetry(work, 10))
	doWork()

	doWork2 := WithRetry(WithLogging(work), 10)
	doWork2()
}

func work() bool {
	if rand.Intn(3) == 0 {
		//fmt.Println("work done")
		return true

	}
	//fmt.Println("work failed")
	return false
}

func WithRetry(fn func() bool, retries int) func() bool {
	return func() bool {
		for i := 0; i < retries; i++ {
			if fn() {
				fmt.Printf("try %d success \n", i)
				return true
			}
			fmt.Printf("try %d failed \n", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("failed after %d retries\n", retries)
		return false
	}
}

func WithLogging(fn func() bool) func() bool {
	return func() bool {
		start := time.Now()
		fmt.Printf("function exec started at %v\n", time.Now())
		a := fn()
		fmt.Printf("function exec ended at %v\n", time.Since(start))
		return a
	}
}
