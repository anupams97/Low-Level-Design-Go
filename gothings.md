# 📘 Go Backend Developer Topics & Why They Matter

## ✅ Core Web Server & API Skills

- **http.Handler, http.HandlerFunc, ServeHTTP**  
  Fundamental HTTP interface in Go; all routing and middleware rely on this.

- **http.ServeMux, http.ListenAndServe**  
  Built-in server/router to handle HTTP requests.

- **Middleware Patterns (e.g., logging, panic recovery)**  
  Add cross-cutting behavior around your handlers cleanly.

- **Request & Response Lifecycle (`http.Request`, `http.ResponseWriter`)**  
  Understand how data flows through a web request.

- **context.Context**  
  Supports timeouts, deadlines, and cancellation for API calls.

- **encoding/json (marshal/unmarshal)**  
  Serialize/deserialize JSON for REST APIs.

- **http.Client, Do, timeouts**  
  Make reliable and timeout-safe outbound API requests.

- **Error handling (wrapping, sentinel errors)**  
  Handle and propagate errors idiomatically and clearly.

---

## 🔐 Security & Reliability

- **recover() for panic handling**  
  Prevent crashes by recovering from panics in middleware.

- **Graceful shutdown (`os.Signal`, `context.WithTimeout`)**  
  Ensure active requests complete before shutting down server.

- **Input validation (e.g., go-playground/validator)**  
  Prevent invalid/untrusted input from reaching business logic.

- **Rate limiting, timeouts, circuit breakers**  
  Build resilient APIs under load or failure.

- **TLS / HTTPS (`http.Server` with TLSConfig)**  
  Secure your APIs with HTTPS in production.

---

## 🧱 Clean Architecture & Structure

- **Handler → Service → Repo layering**  
  Keeps logic separated and testable.

- **Dependency injection (manual, fx, wire)**  
  Makes code loosely coupled and easier to test.

- **Go module layout (internal/, cmd/, pkg/)**  
  Structure large codebases for maintainability.

- **init() vs constructors**  
  Properly initialize services and configs.

---

## 🚀 Performance & Concurrency

- **Goroutines & channels**  
  Lightweight concurrency primitives for async tasks.

- **Worker pools, fan-in/fan-out patterns**  
  Efficient handling of concurrent jobs or messages.

- **sync.Mutex, RWMutex, atomic ops**  
  Safely manage shared state in concurrent code.

- **Benchmarking with `testing.B`**  
  Measure and optimize performance.

- **Profiling with `pprof`**  
  Analyze memory, CPU, and goroutine usage in production.

---

## 📦 Tooling & Common Libraries

- **gorilla/mux, chi**  
  Popular routers with middleware support and path params.

- **zap, zerolog**  
  Structured, performant production logging.

- **gorm, sqlx, pgx**  
  Tools for working with SQL databases.

- **viper, envconfig**  
  Load configs from env vars or files easily.

- **net/http/httptest**  
  Test your HTTP handlers in isolation.

---

## 🧪 Testing

- **Unit tests for handler, service, repo layers**  
  Ensure logic works and regressions are caught early.

- **Table-driven tests**  
  Clean, scalable testing pattern idiomatic to Go.

- **Mocking interfaces (manual or mockgen)**  
  Isolate dependencies in tests.

- **httptest.Server & httptest.NewRequest**  
  Spin up real test HTTP servers and simulate requests.

- **testify/assert, require**  
  Cleaner, more expressive assertions.

---

## 🧰 Advanced Backend Topics

- **gRPC (google.golang.org/grpc)**  
  High-performance APIs with contract-first design.

- **Protobufs**  
  Compact, language-agnostic data format.

- **Kafka, RabbitMQ (sarama, amqp)**  
  Event-driven system support for scale and decoupling.

- **Redis, Aerospike clients**  
  High-speed caching and key-value storage.

- **Rate limiting (token bucket, leaky bucket)**  
  Throttle clients to protect services.

- **JWT auth, session management**  
  Authenticate users securely in stateless services.

- **Swagger / OpenAPI (swaggo)**  
  Auto-generate API documentation for clients and testing tools.

---

# 🎯 Tip: Learn progressively, not all at once.
Start with web basics → add concurrency & testing → expand into architecture & resiliency.