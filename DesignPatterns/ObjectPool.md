```go

package main

import "fmt"

// DatabaseConnection represents a mock database connection
type DatabaseConnection struct {
	ID int
}

// Connect simulates connecting to the database
func (db *DatabaseConnection) Connect() {
	fmt.Printf("DatabaseConnection %d: Connected\n", db.ID)
}

// Disconnect simulates disconnecting from the database
func (db *DatabaseConnection) Disconnect() {
	fmt.Printf("DatabaseConnection %d: Disconnected\n", db.ID)
}

// Validate checks if the connection is still valid
func (db *DatabaseConnection) Validate() bool {
	// Add real validation logic here
	return true
}

package main

import (
"errors"
"fmt"
"sync"
"time"
)

type ObjectPool struct {
	mu        sync.Mutex
	pool      chan *DatabaseConnection
	maxSize   int
	counter   int
	idleTime  time.Duration
	closeChan chan struct{} // Signal to stop background tasks
}

var (
	instance     *ObjectPool
	once         sync.Once
	ErrPoolFull  = errors.New("pool is full")
	ErrPoolEmpty = errors.New("no available objects")
)

// NewObjectPool initializes a new object pool (singleton)
func NewObjectPool(maxSize int, idleTime time.Duration) *ObjectPool {
	once.Do(func() {
		instance = &ObjectPool{
			pool:      make(chan *DatabaseConnection, maxSize),
			maxSize:   maxSize,
			idleTime:  idleTime,
			closeChan: make(chan struct{}),
		}
		// Start the idle timeout handler
		go instance.handleIdleTimeout()
	})
	return instance
}

// Get retrieves an object from the pool or creates a new one if possible
func (p *ObjectPool) Get() (*DatabaseConnection, error) {
	select {
	case conn := <-p.pool:
		// Validate the connection before returning it
		if conn.Validate() {
			return conn, nil
		}
		// If invalid, discard and create a new one
		p.mu.Lock()
		defer p.mu.Unlock()
		p.counter++
		return &DatabaseConnection{ID: p.counter}, nil
	default:
		// Create a new connection if below max size
		p.mu.Lock()
		defer p.mu.Unlock()
		if p.counter < p.maxSize {
			p.counter++
			return &DatabaseConnection{ID: p.counter}, nil
		}
		return nil, ErrPoolEmpty
	}
}

// Put returns an object to the pool
func (p *ObjectPool) Put(conn *DatabaseConnection) error {
	select {
	case p.pool <- conn:
		return nil
	default:
		return ErrPoolFull
	}
}

// handleIdleTimeout checks for idle objects and removes them if necessary
func (p *ObjectPool) handleIdleTimeout() {
	ticker := time.NewTicker(p.idleTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case conn := <-p.pool:
				// Disconnect idle objects
				fmt.Printf("Connection %d: Removed due to idleness\n", conn.ID)
				conn.Disconnect()
			default:
				// No idle objects to remove
			}
		case <-p.closeChan:
			// Stop the handler when the pool is closed
			return
		}
	}
}

// Size returns the current size of the pool
func (p *ObjectPool) Size() int {
	return len(p.pool)
}

// Close shuts down the pool and releases all resources
func (p *ObjectPool) Close() {
	close(p.closeChan)
	for {
		select {
		case conn := <-p.pool:
			conn.Disconnect()
		default:
			return
		}
	}
}


package main

import (
"fmt"
"time"
)

func main() {
	// Create a singleton pool with max size 3 and 5-second idle timeout
	pool := NewObjectPool(3, 5*time.Second)

	// Borrow a connection
	conn1, _ := pool.Get()
	conn1.Connect()

	// Borrow another connection
	conn2, _ := pool.Get()
	conn2.Connect()

	// Return the first connection
	pool.Put(conn1)
	conn1.Disconnect()

	// Borrow a third connection
	conn3, _ := pool.Get()
	conn3.Connect()

	// Display pool size
	fmt.Printf("Pool size: %d\n", pool.Size())

	// Wait for idle timeout to remove idle connections
	time.Sleep(6 * time.Second)

	// Display pool size after idle connections are removed
	fmt.Printf("Pool size after idle timeout: %d\n", pool.Size())

	// Close the pool and release resources
	pool.Close()
}
```

Why Singleton for Object Pool?
1.	Resource Sharing:
•	A single pool ensures that all parts of the application share the same limited resources (e.g., database connections).
2.	Centralized Management:
•	Avoids redundant pools and simplifies configuration.
3.	Thread Safety:
•	sync.Once ensures safe and efficient initialization.

	1.	Singleton Pattern:
	•	Ensures a single global instance of the pool via sync.Once.
	2.	Idle Timeout:
	•	Removes idle objects after a specified duration using a background goroutine.
	3.	Health Checks:
	•	Validates objects before returning them from the pool and discards invalid ones.
	4.	Dynamic Resizing:
	•	Automatically creates new objects if the pool is not full and adjusts dynamically to load.
	5.	Graceful Shutdown:
	•	Close() releases all resources and stops background tasks.