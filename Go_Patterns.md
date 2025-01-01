
# **Go Patterns: Leveraging Functions and Go's Unique Features**

Go’s simplicity, focus on composition, and powerful concurrency model make it ideal for implementing design patterns. This document explores patterns using Go's ability to pass functions as arguments, alongside other Go-specific and traditional design patterns adapted to Go.

---

## **1. Functional Options Pattern**
- **Purpose**: Configures structs or functions using variadic functions.
- **When to Use**: When you need optional or flexible configurations.

### Example
```go
type Server struct {
    Host string
    Port int
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) {
        s.Host = host
    }
}

func WithPort(port int) Option {
    return func(s *Server) {
        s.Port = port
    }
}

func NewServer(options ...Option) *Server {
    server := &Server{Host: "localhost", Port: 8080}
    for _, option := range options {
        option(server)
    }
    return server
}

func main() {
    server := NewServer(WithHost("127.0.0.1"), WithPort(9000))
    fmt.Printf("Server running at %s:%d
", server.Host, server.Port)
}
```

---

## **2. Higher-Order Functions**
- **Purpose**: Functions that take other functions as arguments or return functions for reusable algorithms.
- **When to Use**: For reusable behavior injection.

### Example: Retry Logic
```go
import (
    "errors"
    "fmt"
    "time"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
    for i := 0; i < attempts; i++ {
        if err := fn(); err != nil {
            fmt.Printf("Attempt %d failed: %s
", i+1, err)
            time.Sleep(sleep)
            continue
        }
        return nil
    }
    return errors.New("all attempts failed")
}

func main() {
    err := Retry(3, time.Second, func() error {
        return errors.New("operation failed")
    })
    if err != nil {
        fmt.Println("Operation failed after retries:", err)
    }
}
```

---

## **3. Middleware Pattern**
- **Purpose**: Dynamically chain functions to modify or extend behavior.
- **When to Use**: Commonly used in HTTP servers or message processors.

### Example
```go
import (
    "fmt"
    "net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Request: %s %s
", r.Method, r.URL.Path)
        next(w, r)
    }
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Header.Get("Authorization") == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        next(w, r)
    }
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}

func main() {
    http.Handle("/", LoggingMiddleware(AuthMiddleware(mainHandler)))
    http.ListenAndServe(":8080", nil)
}
```

---

## **4. Callback Pattern**
- **Purpose**: Executes custom behavior dynamically.
- **When to Use**: For event-driven programming or flexible hooks.

### Example
```go
func OnEvent(event string, callback func(string)) {
    fmt.Printf("Event triggered: %s
", event)
    callback(event)
}

func main() {
    OnEvent("UserLogin", func(event string) {
        fmt.Printf("Handling %s event
", event)
    })
}
```

---

## **5. Predicate Functions**
- **Purpose**: Filters or validates data based on criteria.
- **When to Use**: For collection filtering or conditional logic.

### Example
```go
func Filter(numbers []int, predicate func(int) bool) []int {
    result := []int{}
    for _, num := range numbers {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    evenNumbers := Filter(numbers, func(n int) bool { return n%2 == 0 })
    fmt.Println("Even numbers:", evenNumbers)
}
```

---

## **6. Command Pattern**
- **Purpose**: Encapsulates requests as objects to execute dynamically.
- **When to Use**: When actions need to be parameterized or delayed.

### Example
```go
type Command func()

func Execute(commands ...Command) {
    for _, cmd := range commands {
        cmd()
    }
}

func main() {
    printHello := func() { fmt.Println("Hello") }
    printWorld := func() { fmt.Println("World") }

    Execute(printHello, printWorld)
}
```

---

## **7. Event Emitter**
- **Purpose**: Emits events and allows dynamic listeners.
- **When to Use**: For reactive programming or pub-sub mechanisms.

### Example
```go
type EventEmitter struct {
    listeners []func(string)
}

func (e *EventEmitter) On(listener func(string)) {
    e.listeners = append(e.listeners, listener)
}

func (e *EventEmitter) Emit(event string) {
    for _, listener := range e.listeners {
        listener(event)
    }
}

func main() {
    emitter := &EventEmitter{}
    emitter.On(func(event string) { fmt.Println("Listener 1:", event) })
    emitter.On(func(event string) { fmt.Println("Listener 2:", event) })

    emitter.Emit("Event triggered")
}
```

---

## **8. Custom Comparators**
- **Purpose**: Sorts or prioritizes data dynamically.
- **When to Use**: For flexible sorting.

### Example
```go
import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {"Alice", 25},
        {"Bob", 20},
        {"Charlie", 30},
    }

    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println("Sorted by age:", people)
}
```

---

## **9. Retry with Circuit Breaker**
- **Purpose**: Dynamically retries or stops execution based on rules.
- **When to Use**: For error handling and resilience.

### Example
```go
func Retry(attempts int, fn func() error) error {
    for i := 0; i < attempts; i++ {
        if err := fn(); err == nil {
            return nil
        }
    }
    return fmt.Errorf("all attempts failed")
}
```

---

## **10. Data Transformation**
- **Purpose**: Transforms data using a pipeline.
- **When to Use**: For reusable data processing.

### Example
```go
func Map(numbers []int, transformer func(int) int) []int {
    result := []int{}
    for _, num := range numbers {
        result = append(result, transformer(num))
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3}
    squared := Map(numbers, func(n int) int { return n * n })
    fmt.Println("Squared numbers:", squared)
}
```

---

