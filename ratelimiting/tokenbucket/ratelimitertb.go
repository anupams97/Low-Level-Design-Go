package main

import (
	"net/http"
	"sync"
	"time"
)

// Tbucket represents a token Tbucket for rate limiting
type Tbucket struct {
	Capacity   int
	RateLimit  float64
	tokens     float64
	lastUpdate time.Time
	mu         sync.Mutex
}

// NewTbucket creates and initializes a new token Tbucket
func NewTbucket(capacity int, rateLimit float64) *Tbucket {
	return &Tbucket{
		Capacity:   capacity,
		RateLimit:  rateLimit,
		tokens:     float64(capacity),
		lastUpdate: time.Now(),
	}
}

// IsAllowed checks if a request is allowed and deducts a token if allowed
func (b *Tbucket) IsAllowed() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Refill tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(b.lastUpdate).Seconds()
	b.tokens += elapsed * b.RateLimit
	if b.tokens > float64(b.Capacity) {
		b.tokens = float64(b.Capacity)
	}
	b.lastUpdate = now

	// Allow request if tokens are available
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

// RateLimiter holds the token Tbuckets for all users
type RateLimiter struct {
	Tbuckets map[string]*Tbucket
	mu       sync.Mutex
}

// NewRateLimiter initializes a new per-user rate limiter
func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		Tbuckets: make(map[string]*Tbucket),
	}
}

// GetTbucket retrieves or creates a token Tbucket for the given user
func (rl *RateLimiter) GetTbucket(userID string, capacity int, rateLimit float64) *Tbucket {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if Tbucket, exists := rl.Tbuckets[userID]; exists {
		return Tbucket
	}
	Tbucket := NewTbucket(capacity, rateLimit)
	rl.Tbuckets[userID] = Tbucket
	return Tbucket
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Request served"))
//}

func RateLimitMiddleware(rl *RateLimiter, capacity int, rateLimit float64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Identify user (for simplicity, using an HTTP header here)
		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			http.Error(w, "User ID is required", http.StatusBadRequest)
			return
		}

		// Get or create the user's Tbucket
		Tbucket := rl.GetTbucket(userID, capacity, rateLimit)

		// Check if the request is allowed
		if Tbucket.IsAllowed() {
			handler(w, r)
		} else {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
		}
	}
}

func nmain() {
	// Create a rate limiter instance
	rateLimiter := NewRateLimiter()

	http.HandleFunc("/hello", RateLimitMiddleware(rateLimiter, 10, 2)) // 10 capacity, 2 tokens/sec per user

	http.ListenAndServe(":8080", nil)
}
