package main

import (
	"net/http"
	"sync"
	"time"
)

type Bucket struct {
	Capacity  int
	RateLimit float64
	Tokens    int
	stop      chan struct{}
	mu        sync.Mutex
}

func (b *Bucket) Stop() {
	b.mu.Lock()
	defer b.mu.Unlock()
	close(b.stop)
}

func (b *Bucket) IsAllowed() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Tokens > 0 {
		b.Tokens--
		return true
	}
	return false
}

func (b *Bucket) Init() {
	ticker := time.NewTicker(time.Duration(float64(time.Second) / float64(b.RateLimit)))
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				b.Increment()
			case <-b.stop:
				return
			}
		}
	}()
}

func (b *Bucket) Increment() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Tokens < b.Capacity {
		b.Tokens++
	}
}

func NewBucket(capacity int, rl float64) *Bucket {
	b := &Bucket{
		Capacity:  capacity,
		RateLimit: rl,
		Tokens:    capacity, // Start with full capacity
		stop:      make(chan struct{}),
	}
	b.Init()
	return b
}

//func main() {
//	b := NewBucket(10, 10)
//	go func() {
//		i := 0
//		for {
//			if b.IsAllowed() {
//				fmt.Printf("pass %d, %v", i, time.Now().Second())
//				i++
//				fmt.Printf("\n")
//			}
//			time.Sleep(100 * time.Millisecond)
//		}
//	}()
//	time.Sleep(10 * time.Second)
//	b.Stop()
//}

func main() {
	bucket := NewBucket(10, 0.1)

	http.HandleFunc("/hello", TokenBucketMiddleware(handler, bucket))
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Request served"))

}

func TokenBucketMiddleware(next http.HandlerFunc, bucket *Bucket) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bucket.IsAllowed() {
			next(w, r)
		} else {
			http.Error(w, "too many request", http.StatusTooManyRequests)
			return
		}
	}
}
