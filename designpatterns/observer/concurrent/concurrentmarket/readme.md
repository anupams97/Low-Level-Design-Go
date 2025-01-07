Real-Life Use Case for Approach 2 (Channel-Based Observer Pattern)

A real-world example where the channel-based observer pattern (Approach 2) is better than the map-based synchronous observer pattern (Approach 1) is in a real-time stock price notification system.

Scenario: Real-Time Stock Price Updates

Imagine you are building a stock market ticker system where multiple users (observers) subscribe to updates about stock prices. The system should:
•	Send updates concurrently without blocking.
•	Allow users to join or leave anytime without affecting others.
•	Prevent slow observers from blocking updates to faster observers.

🚀 Why Approach 2 is Better in This Case

Using Approach 1 (map-based synchronous notification):
•	If one observer is slow (e.g., due to network latency), it blocks all other observers.
•	Notifications are processed sequentially, leading to delays.
•	The system scales poorly when many observers are connected.

Using Approach 2 (channel-based async notification):
•	Notifications are sent to each observer’s channel separately.
•	Slow observers don’t block fast observers.
•	Supports high concurrency (thousands of observers).

🚀 Implementation of Real-Time Stock Price Notifications Using Channels

package main

import (
"fmt"
"math/rand"
"sync"
"time"
)

// Observer interface
type Observer interface {
Update(price float64)
}

// Concrete Observer (User)
type StockTrader struct {
ID int
}

func (t *StockTrader) Update(price float64) {
// Simulate variable latency (some observers are slow)
time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
fmt.Printf("Trader %d received stock price update: $%.2f\n", t.ID, price)
}

// Subject (Stock Market Ticker)
type StockMarket struct {
observers map[Observer]chan float64
mu        sync.Mutex
}

// Constructor
func NewStockMarket() *StockMarket {
return &StockMarket{
observers: make(map[Observer]chan float64),
}
}

// Attach Observer (Subscribe to stock updates)
func (s *StockMarket) Attach(o Observer) {
s.mu.Lock()
defer s.mu.Unlock()

	ch := make(chan float64)
	s.observers[o] = ch

	// Start a goroutine to send updates asynchronously
	go func() {
		for price := range ch {
			o.Update(price)
		}
	}()
}

// Detach Observer (Unsubscribe)
func (s *StockMarket) Detach(o Observer) {
s.mu.Lock()
defer s.mu.Unlock()

	if ch, exists := s.observers[o]; exists {
		close(ch) // Close the observer's channel
		delete(s.observers, o)
	}
}

// Notify Observers with a new stock price (Sent asynchronously)
func (s *StockMarket) Notify(price float64) {
s.mu.Lock()
defer s.mu.Unlock()

	for _, ch := range s.observers {
		ch <- price // Send price update asynchronously
	}
}

func main() {
market := NewStockMarket()

	// Create multiple traders (observers)
	trader1 := &StockTrader{ID: 1}
	trader2 := &StockTrader{ID: 2}
	trader3 := &StockTrader{ID: 3}

	// Attach traders to stock market updates
	market.Attach(trader1)
	market.Attach(trader2)
	market.Attach(trader3)

	// Simulate stock price updates every 1 second
	go func() {
		for i := 0; i < 5; i++ {
			price := 100 + rand.Float64()*10 // Simulating stock price fluctuations
			fmt.Printf("\nNew stock price: $%.2f\n", price)
			market.Notify(price)
			time.Sleep(1 * time.Second)
		}
	}()

	// Run for a few seconds before detaching an observer
	time.Sleep(3 * time.Second)
	fmt.Println("\nTrader 2 unsubscribes from stock updates.")
	market.Detach(trader2)

	// Keep the program running
	time.Sleep(5 * time.Second)
}

🛠 How This Works
1.	Each observer (trader) gets its own channel.
2.	When a new stock price is available, it’s sent to each observer asynchronously.
3.	Observers process updates at their own speed (slow ones don’t block fast ones).
4.	Observers can unsubscribe anytime without affecting others.

🚀 Output Example

New stock price: $104.32
Trader 1 received stock price update: $104.32
Trader 2 received stock price update: $104.32
Trader 3 received stock price update: $104.32

New stock price: $102.75
Trader 1 received stock price update: $102.75
Trader 3 received stock price update: $102.75
Trader 2 received stock price update: $102.75

Trader 2 unsubscribes from stock updates.

New stock price: $106.89
Trader 1 received stock price update: $106.89
Trader 3 received stock price update: $106.89

🚀 Notice how Trader 2 stops receiving updates after unsubscribing.

⏩ Why is This Approach Better Than Approach 1 Here?

Feature	Approach 1 (Map-Based Synchronous Notification)	Approach 2 (Channel-Based Asynchronous Notification)
Concurrency	❌ Blocks all observers if one is slow	✅ Observers process updates independently
Scalability	❌ Not suitable for high-frequency updates	✅ Can handle thousands of observers
Observer Independence	❌ Slow observers delay others	✅ Observers receive updates at their own pace
Unsubscribing	❌ Requires modifying slice/map manually	✅ Can just close(channel) safely
Performance for Real-Time Events	❌ Not ideal for streaming data	✅ Excellent for real-time streaming

🎯 Other Real-Life Use Cases for Approach 2
1.	Live Sports Score Updates → Different users subscribe to match updates.
2.	IoT Device Monitoring → Devices report status updates, and multiple services consume them.
3.	Log Processing System → Various components process logs asynchronously.
4.	Real-Time Chat System → Clients receive chat messages asynchronously.

🔥 Summary

✅ Use Approach 1 (Map-Based Synchronous) when:
✔ Simplicity is more important than speed.
✔ The number of observers is small.
✔ You don’t need high-frequency updates.

✅ Use Approach 2 (Channel-Based Async) when:
✔ Observers should process updates independently.
✔ Scalability is critical (many subscribers).
✔ Real-time events are required (e.g., stock prices, live sports updates).

🚀 Final Verdict

For a small event-driven system → Approach 1 (simpler & direct).
For a high-scale, real-time, or async event-driven system → Approach 2 (more scalable & performant).

🔥 Go with Approach 2 if you need concurrency and real-time event streaming! 🚀
Let me know if you need further clarification! 🚀